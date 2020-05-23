package dao

import (
	"lawyer/model"
	"lawyer/utils"
)

//GetArticles 获取所有相关文章
func GetArticles() ([]*model.Article, error) {
	sqlStr := "select id,author,authorid,genre,title,text"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var articles []*model.Article
	for rows.Next() {
		article := &model.Article{}
		rows.Scan(&article.ID, &article.Author, &article.Authorid, &article.Genre, &article.Title, &article.Text)
		articles = append(articles, article)
	}
	return articles, nil
}