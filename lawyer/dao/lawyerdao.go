package dao

import (
	"lawyer/model"
	"lawyer/utils"
)

// CheckLawyerNameAndPassword 验证律师账号和密码
func CheckLawyerNameAndPassword(lawyername string, password string) (*model.Lawyer, error) {
	sqlStr := "select id,lawyername,password from lawyer where lawyername = ? and password = ?"
	row := utils.Db.QueryRow(sqlStr, lawyername, password)
	lawyer := &model.Lawyer{}
	row.Scan(&lawyer.ID, &lawyer.Lawyername, &lawyer.Password, &lawyer.Name, &lawyer.Genre, &lawyer.Introduction, &lawyer.Phone)
	return lawyer, nil
}

// CheckLawyerName 查询客户账号
func CheckLawyerName(lawyername string) (*model.Lawyer, error) {
	sqlStr := "select id,lawyername,password from lawyer where lawyername = ?"
	row := utils.Db.QueryRow(sqlStr, lawyername)
	lawyer := &model.Lawyer{}
	row.Scan(&lawyer.ID, &lawyer.Lawyername, &lawyer.Password, &lawyer.Name, &lawyer.Genre, &lawyer.Introduction, &lawyer.Phone)
	return lawyer, nil
}

//SaveLawyer 向数据库中插入用户信息
func SaveLawyer(lawyername string, password string, name string, genre string, introduction string, phone string) error {
	sqlStr := "insert into lawyer(lawyername,password,name,genre,introduction,phone) values(?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, lawyername, password, name, genre, introduction, phone)
	if err != nil {
		return err
	}
	return nil
}