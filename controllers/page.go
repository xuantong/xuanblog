package controllers

import (
	"../models"
)

type PageController struct {
	BaseController
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
