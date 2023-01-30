package main

import (
	"ark-online-excel/dao"
	"ark-online-excel/logger"
	//"context"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	_ "ark-online-excel/routers"
	_ "ark-online-excel/service/ws"

	beegoContext "github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	//"ark-online-excel/utils"
)


func signalsWait() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	for {
		s := <-ch
		switch s {
		case syscall.SIGQUIT:
			logger.ZapLogger.Info("用户发送QUIT字符(Ctrl+/)触发")
			return
		case syscall.SIGTERM:
			logger.ZapLogger.Info("结束程序(可以被捕获、阻塞或忽略)")
			return
		case syscall.SIGINT:
			logger.ZapLogger.Info("用户发送INTR字符(Ctrl+C)触发")
			return
		case syscall.SIGHUP:
			logger.ZapLogger.Info("终端控制进程结束(终端连接断开)")
			return
		default:
			logger.ZapLogger.Info("其他原因结束程序")
			return
		}
	}
}

type ExceptionRet struct {
	Code   int         `json:"code"`
	Status bool        `json:"status"`
	Err    string      `json:"error"`
	Data   interface{} `json:"data"`
}

var Exception = ExceptionRet{
	Code:   500,
	Status: false,
	Err:    "panic",
}

func RecoverPanic(ctx *beegoContext.Context) {
	if err := recover(); err != nil {
		if err == beego.ErrAbort {
			return
		}
		if !beego.BConfig.RecoverPanic {
			panic(err)
		}
		var stack string
		logger.ZapLogger.Error(fmt.Sprintf("the request url is %s", ctx.Input.URL()))
		logger.ZapLogger.Error("Handler crashed with error", zap.Any("err", err))
		for i := 1; ; i++ {
			_, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			stack = stack + fmt.Sprintln(fmt.Sprintf("%s:%d", file, line))

		}
		logger.ZapLogger.Error("panic stack", zap.Any("stack", stack))
		_ = ctx.Output.JSON(Exception, false, false)
	}
}

func main(){

	RunMode := os.Getenv("RunMode")
	if RunMode == "" {
		RunMode = beego.AppConfig.String("RunMode")
	}

	beego.BConfig.RunMode = RunMode
	beego.BConfig.RecoverFunc = RecoverPanic
	//utils.TeamAddr = beego.AppConfig.String("TeamAddr")
	logger.ZapLogger.Info(fmt.Sprintf("运行环境=%s", RunMode))


	if err := dao.MysqlInit(); err != nil {
		beego.Info("init dao.error", zap.Error(err))
		return
	}

	// 跨域
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	// start
	//rootCtx, _ := context.WithCancel(context.Background())


	go func() {
		defer func() {
			if err := recover(); err != nil {
				logger.ZapLogger.Error("rabbitmq goroutine err", zap.Any("panic", err))
			}
		}()
	}()

	beego.Run()
}