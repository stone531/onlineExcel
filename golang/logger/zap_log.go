package logger

import (
	"github.com/astaxie/beego"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

//　初始化zaplogger日志库
var ZapLogger *zap.Logger

func initLogger(logpath string) *zap.Logger {
	//consoleInfo := lumberjack.Logger{
	//	Filename:   logpath + "/console.log", // 日志文件路径
	//	MaxSize:    128,                   // 每个日志文件保存的最大尺寸 单位：M
	//	MaxBackups: 30,                    // 日志文件最多保存多少个备份
	//	MaxAge:     7,                     // 文件最多保存多少天
	//	Compress:   true,                  // 是否压缩
	//}
	hookInfo := lumberjack.Logger{
		Filename:   logpath + "/info.log", // 日志文件路径
		MaxSize:    128,                   // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,                    // 日志文件最多保存多少个备份
		MaxAge:     7,                     // 文件最多保存多少天
		Compress:   true,                  // 是否压缩
	}
	hookWarn := lumberjack.Logger{
		Filename:   logpath + "/warn.log", // 日志文件路径
		MaxSize:    128,                   // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,                    // 日志文件最多保存多少个备份
		MaxAge:     7,                     // 文件最多保存多少天
		Compress:   true,                  // 是否压缩
	}
	hookError := lumberjack.Logger{
		Filename:   logpath + "/error.log", // 日志文件路径
		MaxSize:    128,                    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,                     // 日志文件最多保存多少个备份
		MaxAge:     7,                      // 文件最多保存多少天
		Compress:   true,                   // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 短路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	//// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)

	// 实现两个判断日志等级的interface
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel
	})

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel && lvl < zapcore.ErrorLevel
	})

	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	core := zapcore.NewTee(
		//zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),zapcore.AddSync(&consoleInfo)), atomicLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(&hookInfo), infoLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(&hookWarn), warnLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(&hookError), errorLevel),
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()

	// 设置初始化字段
	filed := zap.Fields(zap.String("serviceName", "yunji"))
	// 构造日志
	logger := zap.New(core, caller, development, filed)

	logger.Info("log 初始化成功")
	return logger
}

// 初始化日志
func init() {
	logPath := beego.AppConfig.String("logPath")
	if logPath == "" {
		logPath = os.Getenv("logPath")
	}
	ZapLogger = initLogger(logPath)
}
