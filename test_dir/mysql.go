package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "xuantong:xuantong@tcp(10.0.0.9:3306)/xuanblog?charset=utf8")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	var title string
	err = db.QueryRow("select title from xuanblog_article limit 1").Scan(&title)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(title)

}
