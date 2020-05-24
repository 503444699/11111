package model

// Question 结构
type Question struct {
	ID int
	Name string
	Username string
	Phone string
	Genre string 
	State string
}

type Questions struct {
	QuestionList []Question
}
