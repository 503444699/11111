package controller

import (
	"net/http"
	"lawyer/model"
	"lawyer/utils"
	"lawyer/dao"
)

// GetQuestions 获取所有的案件
func GetQuestions(w http.ResponseWriter, r *http.Request) {
//调用所有案件
questions, _ := dao.GetQuestions()
t := template.Must(template.ParseFiles("views/pages/manager/question.html"))
t.Execute(w, books)
// }

// AddQuestion 添加案件信息
func AddQuestion(w http.ResponseWriter, r *http.Request) {
	//获取案件信息
	id := r.PostFormValue("id")
	username := r.PostFormValue("username")
	name := r.PostFormValue("name")
	phone := r.PostFormValue("phone")
 	genre := r.PostFormValue("genre")
 	state := r.PostFormValue("state")
	//创建案件
	question := &model.Question{
		ID: int,
		Username: username,
		Name: name,
		Phone: phone,
		Genre: genre,
		State: state,

 	}
}