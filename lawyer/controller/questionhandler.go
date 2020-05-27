package controller

import (
	"lawyer/dao"
	"html/template"
	"net/http"
	"fmt"
)

// GetQuestions 获取所有的案件
func GetQuestions(w http.ResponseWriter, r *http.Request) {
    //调用所有案件
    questions, _ := dao.GetQuestions()
	t := template.Must(template.ParseFiles("views/pages/lawyer/question.html"))
	fmt.Print(questions[0])
    t.Execute(w, questions)
}

// QuestionAdd 添加
func QuestionAdd(w http.ResponseWriter, r *http.Request)  {
	number := r.PostFormValue("number")
	name := r.PostFormValue("name")
	username := r.PostFormValue("username")
	phone := r.PostFormValue("phone")
	genre := r.PostFormValue("genre")
	state := r.PostFormValue("state")
	// 验证
	question, _ := dao.CheckNumberAndName(number, name)
	fmt.Print(question)
	fmt.Println("id:",question.ID)
	fmt.Println("name:",question.Number)
	fmt.Println("name:",&question.Number)
	fmt.Println("password:",question.Name)
	fmt.Println("password:",&question.Name)
	fmt.Println("user:",question.Username)
	fmt.Println("user:",&question.Username)
	fmt.Println("phone:",question.Phone)
	fmt.Println("phone:",&question.Phone)
	fmt.Println("state:",question.State)
	fmt.Println("s:",&question.State)
	fmt.Println("g:",question.Genre)
	fmt.Println("g:",&question.Genre)
	if question.ID > 0{
		t := template.Must(template.ParseFiles("views/pages/user/quessstion_add.html"))
		t.Execute(w, "案件已存在")
	}else{
		dao.AddQuestion(number, name, username, phone, genre, state)
		t := template.Must(template.ParseFiles("views/pages/user/add_success.html"))
		t.Execute(w, "")
	}
}

//CheckNumber 验证
func CheckNumber(w http.ResponseWriter, r *http.Request) {
	number := r.PostFormValue("number")
	question, _ := dao.CheckNumber(number)
	fmt.Print(question)
	if question.ID > 0 {
		w.Write([]byte("案件已存在"))
	} else {
		w.Write([]byte("<font style='color:green'>案件可注册</font>"))
	}
}