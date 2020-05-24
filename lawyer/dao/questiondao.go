package dao

import (
	"lawyer/model"
	"lawyer/utils"
	"fmt"
)

//GetQuestions 获取数据库中所有的图书
func GetQuestions() ([]*model.Question, error) {
	//写sql语句
	sqlStr := "select id,name,username,phone,genre,state from question"
	//执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var questions []*model.Question
	for rows.Next() {
		question := &model.Question{}
		//给question中的字段赋值
		rows.Scan(&question.ID, &question.Name, &question.Username, &question.Phone, &question.Genre, &question.State)
		fmt.Print(question)
		//将question添加到questions中
		questions = append(questions, question)
	}
	fmt.Print(questions[0])
	return questions, nil
}

// AddQuestion 向数据库中插入案件信息
func AddQuestion(question *model.Question) error {
	sqlStr := "insert into question(name,username,phone,genre,state) values(?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, question.Name, question.Username, question.Phone, question.Genre, question.State)
	if err != nil {
		return err
	}
	return nil
}

// GetQuestionByID 查询案件id
func GetQuestionByID(questionID string) (*model.Question, error) {
	sqlStr := "select id,name,username,phone,genre,state from question where id = ?"
	row := utils.Db.QueryRow(sqlStr, questionID)
	question := &model.Question{}
	row.Scan(&question.ID, &question.Name, &question.Username, &question.Phone, &question.Genre, &question.State)
	return question, nil
}
