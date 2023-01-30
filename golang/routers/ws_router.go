package routers

import (
	"ark-online-excel/controller"
	"github.com/astaxie/beego"
)

func wsRouter() beego.LinkNamespace {
	ns := beego.NSNamespace("/ws",
		beego.NSRouter("/test", &controllers.ExcelController{}, "get:JoinExcel"),

		// 创建文档
		beego.NSRouter("/onlineexcel/create", &controllers.SheetController{}, "post:CreateSheetData"),
		// 查找文档 OpenSheetMateData
		beego.NSRouter("/onlineexcel/query", &controllers.SheetController{}, "post:QuerySheetFile"),
		// 打开文档
		beego.NSRouter("/onlineexcel/open", &controllers.SheetController{}, "get:OpenSheetMateData"),
		// 删除文档
		beego.NSRouter("/onlineexcel/delete", &controllers.SheetController{}, "get:DelSheetFile"),

		// 修改文档
		//beego.NSRouter("/onlineexcel/edit", &controllers.SheetController{}, "post:SaveSheetData"),
		//保持当前文档
		//beego.NSRouter("/onlineexcel/save", &controllers.SheetController{}, "post:SaveSheetData"),
		// 获取报表的操作历史
		//beego.NSRouter("/onlineexcel/history/get", &controllers.SheetController{}, "get:GetTableRawData"),
		// 获取报表的原生字段
		//beego.NSRouter("/onlineexcel/rawdatas/get", &controllers.SheetController{}, "get:GetSheetTableMeta"),
		// 获取数据源的所有字段信息
		//beego.NSRouter("/onlineexcel/tablemeta/get", &controllers.SheetController{}, "get:GetSheetHistory"),
		//保存版本
		//beego.NSRouter("/onlineexcel/tablemeta/get", &controllers.SheetController{}, "get:GetSheetHistory"),


		//beego.NSRouter("/diagnosis", &controllers.WsController{}, "get:DiagnosisClusterById"),
		//beego.NSRouter("/diagnosis/list", &controllers.WsController{}, "post:GetCronDiagnosisList"),
		//beego.NSRouter("/diagnosis/detail", &controllers.WsController{}, "get:GetCronDiagnosisDetail"),
	)

	return ns
}
