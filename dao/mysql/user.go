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
		return errors.New("用户已存在")
	}
	return nil
}

func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := "select  user_id,username,password from  user where username = ?"
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return errors.New("用户不存在")
	}
	if err != nil {
		zap.L().Error("查询用户失败", zap.Error(err))
		return nil
	}
	// 检查密码是否正确
	password := encodePassword(oPassword)
	if password != user.Password {
		return errors.New("密码错误")
	}
	return nil
}
