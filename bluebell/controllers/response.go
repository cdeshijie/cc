package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
前端期望得到的信息构成，有利于前端处理后端传来的信息
{
	"code": 10001,  //程序中的错误码
	"msg":  xxx,	//提示的错误信息
	"date":{},		//数据
}

*/

type Response struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Date interface{} `json:"date,omitempty"`
}

func ResponseError(c *gin.Context, code ResCode) {
	rd := &Response{
		Code: code,
		Msg:  code.GetMsg(),
		Date: nil,
	}
	c.JSON(http.StatusOK, rd)
}

func ResponseSuccess(c *gin.Context, date interface{}) {
	rd := &Response{
		Code: CodeSuccess,
		Msg:  CodeSuccess.GetMsg(),
		Date: date,
	}
	c.JSON(http.StatusOK, rd)
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	rd := &Response{
		Code: code,
		Msg:  msg,
		Date: nil,
	}
	c.JSON(http.StatusOK, rd)
}
