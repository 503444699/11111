package controller

import (
	"lawyer/model"
	"lawyer/dao"
	"html/template"
	"net/http"
)

// GetQuestions 获取所有的案件
func GetQuestions(w http.ResponseWriter, r *http.Request) {
    //调用所有案件
    questions, _ := dao.GetQuestions()
    t := template.Must(template.ParseFiles("views/pages/lawyer/question.html"))
    t.Execute(w, questions)
}

// AddQuestion 添加案件信息
func AddQuestion(w http.ResponseWriter, r *http.Request) {
	//获取案件信息
	name := r.PostFormValue("name")
	username := r.PostFormValue("username")
	phone := r.PostFormValue("phone")
 	genre := r.PostFormValue("genre")
 	state := r.PostFormValue("state")
	//创建案件
	question := &model.Question{
		Name: name,
		Username: username,
		Phone: phone,
		Genre: genre,
		State: state,
	}
	dao.AddQuestion(question)
	GetQuestions(w, r)
}

//ToAddBook 去添加图书的页面
func ToAddBook(w http.ResponseWriter, r *http.Request) {
	//获取要更新的图书的id
	questionID := r.FormValue("questionID")
	//调用bookdao中获取图书的函数
	question, _ := dao.GetQuestionByID(questionID)
	if question.ID > 0 {
		//解析模板
		t := template.Must(template.ParseFiles("views/pages/user/question_add.html"))
		//执行
		t.Execute(w, "案件信息已注册")
	} else {
		//解析模板
		t := template.Must(template.ParseFiles("views/pages/user/add_success.html"))
		//执行
		t.Execute(w, "")
	}
}
