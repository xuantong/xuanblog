package controllers

import (
	"../models"
	"github.com/astaxie/beego"
)

type PageController struct {
	beego.Controller
}

func (this *PageController) Get() {
	pageId, _ := this.GetInt(":nid")

	// models.RedirectToError(this., err)

	dbConn := models.DBConn()
	defer dbConn.Close()

	this.Data["pageId"] = pageId
	this.Data["article"], _ = models.ArticleById(dbConn, pageId)

	this.TplNames = "page/page.tpl"
}
