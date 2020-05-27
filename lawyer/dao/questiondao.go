package dao

import (
	"lawyer/model"
	"lawyer/utils"
	"fmt"
)

//GetQuestions 获取数据库中所有的案件
func GetQuestions() ([]*model.Question, error) {
	//写sql语句
	sqlStr := "select id,number,name,username,phone,genre,state from question"
	//执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var questions []*model.Question
	for rows.Next() {
		question := &model.Question{}
		//给question中的字段赋值
		rows.Scan(&question.ID, &question.Number, &question.Name, &question.Username, &question.Phone, &question.Genre, &question.State)
		fmt.Print(question)
		//将question添加到questions中
		questions = append(questions, question)
	}
	fmt.Print(questions[0])
	return questions, nil
}

// AddQuestion 向数据库中插入案件信息
func AddQuestion(number string, name string, username string, phone string, genre string, state string) error {
	sqlStr := "insert into question(number,name,username,phone,genre,state) values(?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, number, name, username, phone, genre, state)
	if err != nil {
		return err
	}
	return nil
}

// CheckNumber 查询
func CheckNumber(number string) (*model.Question, error) {
	sqlStr := "select id,number,name,username from question where number = ?"
	row := utils.Db.QueryRow(sqlStr, number)
	question := &model.Question{}
	row.Scan(&question.ID, &question.Number, &question.Name, &question.Username, &question.Phone,&question.Genre, &question.State)
	fmt.Println(question)
	return question, nil
}

// CheckNumberAndName 验证
func CheckNumberAndName(number string, name string) (*model.Question, error) {
	sqlStr := "select id,number,name from question where number = ? and name = ?"
	question := &model.Question{}
	row := utils.Db.QueryRow(sqlStr, number, name)
	row.Scan(&question.ID, &question.Number, &question.Name)
	fmt.Println(question)
	return question, nil
}

