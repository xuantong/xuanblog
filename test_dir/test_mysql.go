package main

import (
	// "database/sql"
	_ "github.com/go-sql-driver/mysql"
	//	"errors"
	"fmt"
	"github.com/coocood/qbs"
	"time"
)

type xuanblog_article struct { // 文章
	Id           int64     // 文章ID
	Title        string    // 文章标题
	Abstract     string    // 文章摘要，Markdown格式
	Abstracthtml string    // 文章摘要，HTML格式
	Content      string    // 文章内容，Markdown格式
	Contenthtml  string    // 文章内容，HTML格式
	Pubdate      time.Time // 发布日期
	Updated      time.Time // 最后更新日期
	Category     int64     // 文章分类(外键)
}

func getArticleList1() (count int) {
	return 1234
}

// 获得文章总数
func getArticleList(q *qbs.Qbs) ([]*xuanblog_article, error) {
	var article []*xuanblog_article
	err := q.Limit(100).Offset(0).FindAll(&article)
	return article, err
}

func main() {
	// db, err := sql.Open("mysql", "xuantong:xuantong@tcp(10.0.0.9:3306)/xuanblog")
	qbs.Register("mysql", "xuantong:xuantong@tcp(10.0.0.9:3306)/xuanblog?parseTime=true", "xuanblog", qbs.NewMysql())

	qbs_conn, _ := qbs.GetQbs()
	qbs_mg, _ := qbs.GetMigration()
	qbs_mg.CreateTableIfNotExists(new(xuanblog_article))

	var article = new(xuanblog_article)
	article.Pubdate = time.Now()
	article.Abstract = "test"
	article.Abstracthtml = "test"
	article.Category = 2
	article.Content = "test"
	article.Contenthtml = "test"
	article.Title = "test"
	article.Updated = time.Now()
	qbs_conn.Save(article)

	article1, _ := getArticleList(qbs_conn)
	fmt.Println(article1[0].Title)
}
