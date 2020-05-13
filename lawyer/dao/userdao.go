package dao

import (
	"lawyer/model"
	"lawyer/utils"
)

// CheckUserNameAndPassword 验证客户账号和密码
func CheckUserNameAndPassword(username string, password string) (*model.User, error) {
	sqlStr := "select id,lawyername,password from users where username = ? and password = ?"
	row := utils.Db.QueryRow(sqlStr, username, password)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Name)
	return user, nil
}

// CheckUserName  查询客户账号
func CheckUserName(username string) (*model.User, error) {
	sqlStr := "select id,username,password from users where username = ?"
	row := utils.Db.QueryRow(sqlStr, username)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Name)
	return user, nil
}
 
// SaveUser 向数据库中插入用户信息
func SaveUser(username string, password string, name string) error {
	sqlStr := "insert into users(username,password,name) values(?,?,?)"
	_, err := utils.Db.Exec(sqlStr, username, password, name)
	if err != nil {
		return err
	}
	return nil
}