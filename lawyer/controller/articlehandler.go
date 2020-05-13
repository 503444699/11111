package controller

import (
	"net/http"
	"lawyer/model"
	"lawyer/utils"
	"lawyer/dao"
)

// GetArticle 获取文章
func GetArticle(w http.ResponseWriter, r *http.Request) {
article, _ := dao.GetArticle()
t := template.Must(template.ParseFiles("views/pages/article/article.html"))
t.Execute(w, books)
}