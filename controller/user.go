package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
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
			c.JSON(http.StatusOK, gin.H{
				"msg": errs.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"code": -1,
			"msg":  removeTopStruct(errs.Translate(trans)),
		})
		return
	}

	fmt.Println(p.Username, p.Password, p.RePassword)
	// 业务处理
	if err := logic.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "注册失败",
		})
		return
	}
	// 返回结果
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "注册成功",
	})

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
			c.JSON(http.StatusOK, gin.H{
				"msg": errs.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"code": -1,
			"msg":  removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	// 业务处理
	if err := logic.Login(p); err != nil {
		zap.L().Error("logic.login failed", zap.String("username", p.Username), zap.Error(err))
		c.JSON(
			http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户名密码失败",
			})
		return

	}

	// 返回响应
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "登录成功",
	})

}
