package dao

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("测试question")
	m.Run()
}

func TestQuestion(t *testing.T) {
	fmt.Println("测试questiondao中的方法")
	t.Run("测试question相关函数",TestGetQuestions)
}
func TestGetQuestions(t *testing.T) {
	questions, _ := GetQuestions()
	for k, v := range questions {
		fmt.Printf("第%v数据为：%v\n", k+1, v)
	}
}