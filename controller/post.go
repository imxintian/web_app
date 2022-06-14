package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"web_app/logic"
	"web_app/models"
)

func CreatePostHandler(c *gin.Context) {
	// 获取参数及参数校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("ShouldBindJSON with invalid param", zap.Any("err", err))
		zap.L().Error("CreatePost with invalid param", zap.Error(err))
		// 判断err是否是validator.ValidationErrors类型
		ResponseError(c, CodeInvalidParam)
		return

	}
	// 从context中获取userId
	userID, err := getCurrentUserId(c)
	if err != nil {
		zap.L().Debug("getCurrentUserId with invalid param", zap.Any("err", err))
		zap.L().Error("getCurrentUserId error", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID

	// 构建新的post
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 返回结果
	ResponseSuccess(c, nil)

}
