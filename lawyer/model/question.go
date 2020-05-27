package model

// Question 结构
type Question struct {
	ID int
	Number string
	Name string
	Username string
	Phone string
	Genre string 
	State string
}

// Questions 结构
type Questions struct {
	QuestionList []Question
}
