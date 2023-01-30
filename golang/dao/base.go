package dao

import (
	"ark-online-excel/logger"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func MysqlInit() error {
	return mysqlInit()
}

// connect to mysql
func mysqlInit() error {
	maxIdle := 10
	maxConn := 50

	dbMysql := beego.AppConfig.String("DbMysql")
	err := orm.RegisterDataBase("default", "mysql", dbMysql, maxIdle, maxConn)
	orm.Debug = true
	if err != nil {
		logger.ZapLogger.Error("mysqlInit", zap.Error(err))
		return err
	} else {
		logger.ZapLogger.Debug("mysql info", zap.String("connection", dbMysql))
	}
	logger.ZapLogger.Info("connect mysql, ok")
	return nil
}
