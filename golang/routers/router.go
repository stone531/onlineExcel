package routers

import (
	"ark-online-excel/controller"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api/v1",
		wsRouter(),
	)
	other := beego.NewNamespace("/api/v1",
		// health
		beego.NSRouter("/health", &controllers.HealthController{}, "get:Health"),
		// current goroutine
		beego.NSRouter("/goroutine", &controllers.HealthController{}, "get:GoRoutine"),
	)
	beego.AddNamespace(ns)
	beego.AddNamespace(other)
}
