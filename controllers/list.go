package controllers

import (
	"../models"
	// "fmt"
	// "github.com/astaxie/beego"
	"github.com/coocood/qbs"
)

type ListController struct {
	BaseController
}

func (this *ListController) Get() {

	this.TplNames = "list/list.tpl"
	// this.Data["count"] = 12132141243d

	// dbConn := models.DBConn()
	// defer dbConn.Close()

	qbs.WithQbs(func(q *qbs.Qbs) error {
		this.Data["articleList"], _ = models.ArticleList(q, 0)
		this.Data["count"] = models.ArticleCount(q)

		return nil
	})

	//this.Data["articleList"] = "xuantong abc "

}
