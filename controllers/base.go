package controllers

import (
	"github.com/astaxie/beego"
)

// 约定：如果子controller 存在NestPrepare()方法，就实现了该接口，
type NestPrepare interface {
	NestPrepare()
}

type BaseController struct {
	beego.Controller
}

func (ctx *BaseController) Prepare() {
	// 将页面路径 保存到 Path变量里面
	ctx.Data["Path"] = ctx.Ctx.Request.RequestURI
	if app, ok := ctx.AppController.(NestPrepare); ok {
		app.NestPrepare()
	}
}