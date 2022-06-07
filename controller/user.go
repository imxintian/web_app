package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/logic"
	"web_app/models"
)

func SignUpHandler(c *gin.Context) {
	// 参数校验
	p := new(models.ParamSignUp)
	// shouldBindJSON() 可以接收任何类型的数据，绑定到结构体中
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断err是否是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, codeInvalidPassword)
			return
		}

		ResponseErrorWithMsg(c, codeInvalidPassword, removeTopStruct(errs.Translate(trans)))
		return
	}

	fmt.Println(p.Username, p.Password, p.RePassword)
	// 业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic signUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回结果
	ResponseSuccess(c, nil)

}

func LoginHandler(c *gin.Context) {
	// 参数校验
	p := new(models.ParamLogin)
	// shouldBindJSON() 可以接收任何类型的数据，绑定到结构体中
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断err是否是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}

		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 业务处理
	if err := logic.Login(p); err != nil {
		zap.L().Error("logic.login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrPasswordIncorrect) {
			ResponseError(c, CodeUserPasswordError)
			return
		}
		ResponseError(c, CodeServerBusy)
		return

	}

	// 返回响应
	ResponseSuccess(c, nil)

}
