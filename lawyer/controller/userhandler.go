package controller

import (
	"lawyer/dao"
	"html/template"
	"net/http"
	"fmt"
)

// UserLogin 客户登陆函数
func UserLogin(w http.ResponseWriter, r *http.Request)  {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	// 验证客户账号和密码
	user, _ := dao.CheckUserNameAndPassword(username, password)
	fmt.Print(username,password)
	if user.ID > 0{
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, "")
	}else{
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "客户账号或密码不正确")
	}
}

// UserRegist 处理客户注册函数
func UserRegist(w http.ResponseWriter, r *http.Request)  {
	// 获取客户账号和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	name := r.PostFormValue("name")
	// 验证客户账号和密码
	user, _ := dao.CheckUserNameAndPassword(username, password)
	fmt.Print(username,password)
	if user.ID > 0{
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "客户账号已存在")
	}else{
		dao.SaveUser(username, password, name)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}
}

//CheckUserName 通过发送Ajax验证用户名是否可用
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	user, _ := dao.CheckUserName(username)
	fmt.Print(username)
	if user.ID > 0 {
		w.Write([]byte("用户名已存在！"))
	} else {
		w.Write([]byte("<font style='color:green'>用户名可用！</font>"))
	}
}