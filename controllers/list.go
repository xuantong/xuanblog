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

	dbConn := models.DBConn()
	defer dbConn.Close()

	this.Data["articleList"], _ = models.ArticleList(dbConn, 0)
	// int64 total = models.ArticleCount(dbConn)

	qbs.WithQbs(func(q *qbs.Qbs) error {
		this.Data["articleList"], _ = models.ArticleList(q, 0)

		var total int64
		total = models.ArticleCount(q)

		// 获取分页导航HTML代码
		this.SetPaginator(10, total)
		// this.Data["Paginator"] = paginator
		this.Data["count"] = total
		return nil
	})

	//this.Data["articleList"] = "xuantong abc "

}
