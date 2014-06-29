package controllers

import (
	"../models"
	// "fmt"
	"github.com/astaxie/beego"
	// "github.com/coocood/qbs"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	this.Layout = "layout/lay_out_list.html"
	this.Data["appname"] = beego.AppConfig.String("appname")
}

func (this *BaseController) SetPaginator(per int, nums int64) *models.Paginator {
	p := models.NewPaginator(this.Ctx.Request, per, nums)
	this.Data["paginator"] = p
	return p
}
