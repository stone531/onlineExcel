package timer

import (
	"ark-online-excel/logger"
	"context"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"

	"ark-online-excel/service/ws"
)

//var DiagnosisCronJobParser cron.Parser

var g_cron cron.Cron


func InitTimer(ctx context.Context) {

	var err error

	DiagnosisCronJobParser := cron.NewParser(
		cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow,
	)
	g_cron := cron.New(cron.WithParser(DiagnosisCronJobParser))
	//添加定时任务 schedule 为cron表达式

	//增加新的定时任务
	esSchedule := "0 */10 * * * *"
	_, err = g_cron.AddFunc(esSchedule, ws.ClearUnUsingShareRoom)
	if err != nil {
		logger.ZapLogger.Error("LoadAllEsTemplate crontab init err", zap.Any("panic", err))
		return
	}

	g_cron.Start()
	//defer g_cron.Stop()

}
