package main
 
import (
	"lawyer/controller"
    "net/http"
    "text/template"
)

// IndexHandler 律师服务系统首页
func IndexHandler (w http.ResponseWriter, r *http.Request) {
    // 解析
    t := template.Must(template.ParseFiles("views/index.html"))
    // 执行
    t.Execute(w, "")
}
 
func main () {
    // 处理静态资源
    http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("views/static"))))
    http.HandleFunc("/main", IndexHandler)
    // 客户登陆
    http.HandleFunc("/pages/user/login",controller.Login)
    // 客户注册
    http.HandleFunc("/pages/user/regist",controller.Regist)
    // 律师注册
    http.HandleFunc("/pages/lawyer/login",controller.Login)
    // 律师登陆
    http.HandleFunc("/pages/lawyer/regist",controller.Regist)
    http.ListenAndServe(":8080", nil)
}
 