package controller

import (
	"lawyer/dao"
	"lawyer/model"
	"lawyer/utils"
	"html/template"
	"net/http"
)

// Login 客户登陆函数
func Login(w http.ResponseWriter, r *http.Request)  {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	// 验证客户账号和密码
	user, _ := dao.CheckUserNameAndPassword(username, password)
	if user.ID > 0{
		t := template.Must(template.ParseFiles("views/pages/user/login_sucess.html"))
		t.Execute(w, "")
	}else{
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "")
	}
}

// Regist 处理客户注册函数
func Regist(w http.ResponseWriter, r *http.Request)  {
	// 获取客户账号和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	name := r.PostFormValue("name")
	// 验证客户账号和密码
	user, _ := dao.CheckUserNameAndPassword(username, password)
	if user.ID > 0{
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已存在")
	}else{
		dao.SaveUser(ausername, password, name)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}
}