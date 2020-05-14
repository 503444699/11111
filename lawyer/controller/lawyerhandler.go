package controller

import (
	"lawyer/dao"
	"html/template"
	"net/http"
)

// LawyerLogin 律师登陆函数
func LawyerLogin(w http.ResponseWriter, r *http.Request)  {
	lawyername := r.PostFormValue("lawyername")
	password := r.PostFormValue("password")
	// 验证客户账号和密码
	user, _ := dao.CheckLawyerNameAndPassword(lawyername, password)
	if user.ID > 0{
		t := template.Must(template.ParseFiles("views/pages/lawyer/login_sucess.html"))
		t.Execute(w, "")
	}else{
		t := template.Must(template.ParseFiles("views/pages/lawyer/login.html"))
		t.Execute(w, "")
	}
}

// LawyerRegist 处理律师注册函数
func LawyerRegist(w http.ResponseWriter, r *http.Request)  {
	lawyername := r.PostFormValue("lawyername")
	password := r.PostFormValue("password")
	name := r.PostFormValue("name")
	genre := r.PostFormValue("genre")
	introduction := r.PostFormValue("introduction")
	phone := r.PostFormValue("phone")
	// 验证客户账号和密码
	user, _ := dao.CheckLawyerNameAndPassword(lawyername, password)
	if user.ID > 0{
		t := template.Must(template.ParseFiles("views/pages/lawyer/regist.html"))
		t.Execute(w, "用户名已存在")
	}else{
		dao.SaveLawyer(lawyername, password, name, genre, introduction, phone)
		t := template.Must(template.ParseFiles("views/pages/lawyer/regist_success.html"))
		t.Execute(w, "")
	}
}
//CheckLawyerName 通过发送Ajax验证用户名是否可用
func CheckLawyerName(w http.ResponseWriter, r *http.Request) {
	lawyername := r.PostFormValue("lawyername")
	lawyer, _ := dao.CheckLawyerName(lawyername)
	if lawyer.ID > 0 {
		w.Write([]byte("用户名已存在！"))
	} else {
		w.Write([]byte("<font style='color:green'>用户名可用！</font>"))
	}
}