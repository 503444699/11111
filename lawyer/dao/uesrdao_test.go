package dao

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	// fmt.Println("测试bookdao中的方法")
	m.Run()
}

func TestUser(t *testing.T) {
	// fmt.Println("测试userdao中的函数")
	// t.Run("验证用户名或密码：", testLogin)
	// t.Run("验证用户名：", testRegist)
	// t.Run("保存用户：", testSave)
}

func testUserLogin(t *testing.T) {
	user, _ := CheckUserNameAndPassword("user5", "123456")
	fmt.Println("获取用户信息是：", user)
}
func testUserRegist(t *testing.T) {
	user, _ := CheckUserName("user5")
	fmt.Println("获取用户信息是：", user)
}
func testSave(t *testing.T) {
	SaveUser("user5", "123456", "张三")
}