package models

import (
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/coocood/qbs"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/go-sql-driver/mysql"
)

func InitDb() {

	database := beego.AppConfig.String("mysqlurls")
	username := beego.AppConfig.String("mysqluser")
	password := beego.AppConfig.String("mysqlpass")

	// fmt.Println(username + ":" + password + "@" + database)
	qbs.Register("mysql", username+":"+password+"@"+database, "xuanblog", qbs.NewMysql())
	// qbs.Register("mysql", "xuantong:xuantong@tcp(10.0.0.9:3306)/xuanblog?parseTime=true", "xuanblog", qbs.NewMysql())
}

func DBConn() (q *qbs.Qbs) {
	qbs_con, err := qbs.GetQbs()
	Check(err)
	return qbs_con
}
