package models

import (
	"github.com/coocood/qbs"
	"time"
)

type xuanblog_article struct { // 文章
	Id       int64     // 文章ID
	Title    string    // 文章标题
	Abstract string    // 文章摘要，HTML格式
	Content  string    // 文章内容，Markdown格式
	Category int       // 文章分类(外键)
	Type     int       // 文本类型,0为html,1为markdown
	Pubdate  time.Time // 发布日期
	Updated  time.Time // 最后更新日期
}

// 获得文章列表
func ArticleList(q *qbs.Qbs, offset int) ([]*xuanblog_article, error) {
	var article []*xuanblog_article
	err := q.Limit(PAGE_SIZE).Offset(0).FindAll(&article)
	return article, err
}

func ArticleCount(q *qbs.Qbs) (count int64) {
	return q.Count("xuanblog_article")
}

func ArticleById(q *qbs.Qbs, id int64) (*xuanblog_article, error) {
	article := new(xuanblog_article)
	article.Id = id
	err := q.Limit(1).Find(article)
	return article, err
}
