package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// TODO:判断用户名是否存在

	if error := mysql.CheckUserExist(p.Username); error != nil {
		return error
	}

	// TODO:生成uid
	userId := snowflake.GenID()
	// 构造一个user 实例
	user := &models.User{
		UserId:   userId,
		Username: p.Username,
		Password: p.Password,
	}
	// TODO:数据入库
	return mysql.InsertUser(user)
}
