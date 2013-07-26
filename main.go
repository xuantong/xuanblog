package main

import (
	"./controllers"
	"./models"
	"github.com/astaxie/beego"
)

func main() {
	models.InitDb()
	beego.Router("/", &controllers.ListController{})
	beego.Router("/list", &controllers.ListController{})
	beego.Router("/p/:nid([0-9]+)", &controllers.PageController{})

	beego.SessionOn = true
	beego.Run()
}
