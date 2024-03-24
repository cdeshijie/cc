package controllers

import (
	"gin_demo2/logic"
	"gin_demo2/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CreatePostHandler 创建帖子
func CreatePostHandler(c *gin.Context) {
	//1.获取并校验参数,这里通过validator的标签binding判断了参数，所以不用自己校验了
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c.ShouldBindJSON post failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//从c中获取userid
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	//2.转入逻辑层，创建帖子,就是把帖子数据存入数据库
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回处理信息，创建成功或者失败
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 根据帖子id得到帖子详情
func GetPostDetailHandler(c *gin.Context) {
	//1.校验绑定参数,获得帖子id
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("Get PostDetail with InvalidParam", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//2.转入logic层,根据id获取帖子数据
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostById failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回结果
	ResponseSuccess(c, data)
}

// GetPostListHandler 获取帖子列表的接口
func GetPostListHandler(c *gin.Context) {
	//1.校验参数,本次获取分页参数
	page, size := getPageInfo(c)
	//2.转入logic层实现
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回结果
	ResponseSuccess(c, data)
}

// GetPostListHandler2 升级版帖子列表接口，根据前端传来的参数动态的获取帖子列表，比如按分数或者创建时间获取列表
// 1.校验参数
// 2.去redis查询帖子id列表
// 3.根据id去数据库查询帖子详情
func GetPostListHandler2(c *gin.Context) {
	//GET请求：/api/v1/posts2?page=1&size=10&order=time
	//1.校验参数,本次获取分页参数
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//2.转入logic层实现
	data, err := logic.GetPostList2(p)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回结果
	ResponseSuccess(c, data)
}

// GetCommunityPostListHandler 根据社区去查询帖子列表,指定了查询数据量的大小（page，size），并且是有序的，用社区表和排序表取交集
func GetCommunityPostListHandler(c *gin.Context) {
	//1.校验参数,本次获取分页参数
	p := &models.ParamCommunityPostList{
		ParamPostList: &models.ParamPostList{
			Page:  1,
			Size:  10,
			Order: models.OrderTime,
		},
	}
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("GetCommunityPostListHandler with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//2.转入logic层实现
	data, err := logic.GetCommunityPostList(p)
	if err != nil {
		zap.L().Error("logic.GetCommunityPostList2 failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回结果
	ResponseSuccess(c, data)
}
