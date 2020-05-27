package model

// Article 结构
type Article struct {
	ID int
	Author string
	Authorid int
	Genre string
	Title string
	Text string
}

// Articles 结构
type Articles struct {
	ArticleList []Article
}