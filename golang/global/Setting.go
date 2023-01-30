package global

import (
	"github.com/astaxie/beego"
)

/**
 全局配置文件
 */

func GetBaseFilePath() string {
	return beego.AppConfig.String("ExcelFileDir")
}

func GetFileUrl() string {
	return beego.AppConfig.String("ReleaseUrl")
}