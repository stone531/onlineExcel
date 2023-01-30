package controllers

import (
	"runtime/pprof"
)

type HealthController struct {
	BaseController
}

func (h *HealthController) Health() {
	h.SetData("ok")
}

func (h *HealthController) GoRoutine() {
	p := pprof.Lookup("goroutine")
	p.WriteTo(h.Ctx.ResponseWriter, 1)
	h.SetData("ok")
}
