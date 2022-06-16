package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
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

// GetPostDetailHandler  获取帖子详情
func GetPostDetailHandler(c *gin.Context) {
	// 获取参数及参数校验
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Debug("parse communityID error", zap.Error(err))
		zap.L().Error("getPostList with invalid param", zap.Error(err))
		return
	}
	// 查询post列表
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("getPostList error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, data)
}

// GetPostListHandler 获取帖子列表
func GetPostListHandler(c *gin.Context) {
	// 获取参数及参数校验
	page, pageSize := getPageInfo(c)
	// 查询post列表
	data, err := logic.GetPostList(page, pageSize)
	if err != nil {
		zap.L().Error("getPostList error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, data)
}
