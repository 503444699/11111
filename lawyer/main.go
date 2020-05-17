package main
 
import (
	"lawyer/controller"
    "net/http"
    "html/template"
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
    // 去html页面
    http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
    // 去首页
    http.HandleFunc("/main", IndexHandler)
    // 客户登陆
    http.HandleFunc("/views/pages/user/login",controller.UserLogin)
    // 客户注册
    http.HandleFunc("/views/pages/user/regist",controller.UserRegist)
    // 律师登陆
    http.HandleFunc("/views/pages/lawyer/login",controller.LawyerLogin)  
    // 律师注册
    http.HandleFunc("/views/pages/lawyer/regist",controller.LawyerRegist)
    // 查看相关文档
    http.HandleFunc("/getArticle",controller.GetArticles)
    // 查看案件信息
    http.HandleFunc("/getQuestions", controller.GetQuestions)
    // 添加案件信息
    http.HandleFunc("/addQuestions", controller.AddQuestion)

    http.ListenAndServe(":8080",nil)
}
 