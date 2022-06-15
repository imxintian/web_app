package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"go.uber.org/zap"
	"web_app/models"
)

const secret = "wxt"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrPasswordIncorrect = errors.New("password incorrect")
	ErrUserExist         = errors.New("user exist")
)

func QueryUserByUserName() {
}

// InsertUser insert a new User into database and returns
func InsertUser(user *models.User) (err error) {
	user.Password = encodePassword(user.Password)
	sqlStr := "insert into user(user_id,username,password) values(?,?,?)"
	_, err = db.Exec(sqlStr, user.UserId, user.Username, user.Password)
	if err != nil {
		panic(err)
	}
	return
}
func encodePassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

// CheckUserExist 检查指定用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := "select count(*) from user where username = ?"
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrUserExist
	}
	return nil
}

func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := "select  user_id,username,password from  user where username = ?"
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrUserNotFound
	}
	if err != nil {
		zap.L().Error("查询用户失败", zap.Error(err))
		return nil
	}
	// 检查密码是否正确
	password := encodePassword(oPassword)
	if password != user.Password {
		return ErrPasswordIncorrect
	}
	return nil
}

//GetUserById 查询用户信息
func GetUserById(id int64) (*models.User, error) {
	sqlStr := "select user_id,username,password from user where user_id = ?"
	user := &models.User{}
	err := db.Get(user, sqlStr, id)
	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}
