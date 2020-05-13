package controller

import (
	"html/template"
	"net/http"
	"lawyer/dao"
)

// GetArticles 获取文章
func GetArticles(w http.ResponseWriter, r *http.Request) {
    articles, _ := dao.GetArticles()
    t := template.Must(template.ParseFiles("views/pages/article/article.html"))
    t.Execute(w, articles)
}