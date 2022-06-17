package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/logic"
	"web_app/models"
)

// PostVoteHandler

func PostVoteHandler(c *gin.Context) {
	// 1. 获取参数
	p := new(models.PostVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Vote with invalid param", zap.Error(err))
		// 判断err是否是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return

	}
	// 获取当前用户Id
	userID, err := getCurrentUserId(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Debug("VoteForPost", zap.Int64("userID", userID), zap.Any("p", p))
		zap.L().Error("VoteForPost error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
