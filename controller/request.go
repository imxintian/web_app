package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

var ErrorUserNotLogin = errors.New("user not login")

const CtxUserIDKey = "userID"

//getCurrentUser 获取当前用户
func getCurrentUserId(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		return 0, ErrorUserNotLogin

	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
	}
	return
}

// getPageInfo 获取分页信息
func getPageInfo(c *gin.Context) (page, pageSize int) {
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")
	page, _ = strconv.Atoi(pageStr)
	pageSize, _ = strconv.Atoi(pageSizeStr)
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 3
	}
	return

}
