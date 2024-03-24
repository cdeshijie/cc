package controllers

import (
	"gin_demo2/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CommunityHandler 获得所有社区
func CommunityHandler(c *gin.Context) {
	//1.接受和校验参数，这里是用户登录后看到的页面，所以应该看到的是所有的社区（community_id,community_name）
	//2.转入logic层处理业务，因为这里是get没有参数，所以无需校验参数
	date, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) //详细错误不对外暴露，只说系统繁忙，除非是提示类
		return
	}
	//3.返回响应
	ResponseSuccess(c, date)
}

// CommunityDetailHandler 根据社区id获得社区详情
func CommunityDetailHandler(c *gin.Context) {
	//1.获取社区的id
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//2.转入logic
	date, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) //详细错误不对外暴露，只说系统繁忙，除非是提示类
		return
	}
	//3.返回响应
	ResponseSuccess(c, date)
}
