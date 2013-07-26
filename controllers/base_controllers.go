package controllers

import (
	// "../models"
	// "fmt"
	"github.com/astaxie/beego"
	// "github.com/coocood/qbs"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	this.Layout = "layout/lay_out1.html"
	this.Data["appname"] = beego.AppConfig.String("appname")
}
