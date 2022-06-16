package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
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

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递的是指针
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	// 生成jwt
	token, err := jwt.GenToken(user.UserId, user.Username)
	if err != nil {
		return nil, err
	}
	user.Token = token
	return user, nil
}

// GetUserById 获取用户信息
func GetUserById(id int64) (*models.User, error) {
	user, err := mysql.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
