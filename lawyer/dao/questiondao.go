package dao

import (
	"lawyer/model"
	"lawyer/utils"
)

//GetQuestions 获取所有案件信息
func GetQuestions() ([]*model.Question, error) {
	sqlStr := "select id,username,name,phone,genre,state"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var questions []*model.Question
	for rows.Next() {
		question := &model.Question{}
		rows.Scan(&question.ID, &question.Username, &question.Name, &question.Phone, &question.Genre, &question.State)
		questions = append(questions, question)
	}
	return questions, nil
}

//AddQuestion 添加案件
func AddQuestion(b *model.Question) error {
	slqStr := "insert into books(username,name,phone,genre,state) values(?,?,?,?,?)"
	_, err := utils.Db.Exec(slqStr, username, name, phone, genre, state)
	if err != nil {
		return err
	}
	return nil
}