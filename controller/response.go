package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	{
		"code": 200,
		"message": "OK",
		"data": {
			"id": 1,
			"name": "test",
			"age": 18
		}
*/

type Response struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResponseError(c *gin.Context, code ResCode) {
	rd := &Response{
		Code: code,
		Msg:  Msg(code),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	rd := &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}
func ResponseSuccess(c *gin.Context, data interface{}) {
	rd := &Response{
		Code: CodeSuccess,
		Msg:  Msg(CodeSuccess),
		Data: data,
	}
	c.JSON(http.StatusOK, rd)
}
