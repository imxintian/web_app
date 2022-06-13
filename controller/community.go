package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/logic"
)

func CommunityHandler(c *gin.Context) {
	// 查询到所有社区（community_id community_name）以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("getCommunityList error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler 获取社区详情
func CommunityDetailHandler(c *gin.Context) {
	// 查询到指定社区的详细信息（community_id community_name）
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("parse id error", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("getCommunityDetail error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
