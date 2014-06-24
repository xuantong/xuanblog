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
	page, _ := this.GetInt("p")
	// this.Data["count"] = 12132141243d

	// dbConn := models.DBConn()
	// defer dbConn.Close()

	qbs.WithQbs(func(q *qbs.Qbs) error {
		this.Data["articleList"], _ = models.ArticleList(q, 0)
		total = models.ArticleCount(q)

		// 获取分页导航HTML代码
		paginator := GetPaginator(total, ItemsPerPage, pagenum)
		this.Data["Paginator"] = paginator
		return nil
	})

	//this.Data["articleList"] = "xuantong abc "

}
