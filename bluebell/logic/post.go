package logic

import (
	"gin_demo2/dao/mysql"
	"gin_demo2/dao/redis"
	"gin_demo2/models"
	"gin_demo2/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	//1.生成post的id通过雪花算法
	p.ID = snowflake.GetID()
	//2.保存到数据库
	err = mysql.CreatePost(p) //mysql.CreatePost(p)的返回值已经是一个错误了，所以直接返回函数结果就行
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.ID, p.CommunityID)
	return
	//3.返回结果,直接返回第二步函数即可
}

// GetPostById 根据帖子id查询帖子详情
func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {
	//查询并组合接口想用的数据
	//获得帖子信息
	post, err := mysql.GetPostById(pid)
	if err != nil {
		return
	}
	//根据作者id获取作者信息
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserById failed",
			zap.Int64("author_id", post.CommunityID),
			zap.Error(err))
		return
	}
	//根据社区id获得社区信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID failed",
			zap.Int64("community_id", post.CommunityID),
			zap.Error(err))
		return
	}
	//将信息写入data
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return
}

// GetPostList 获取帖子列表
func GetPostList(page int64, size int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		//根据作者id获取作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById failed",
				zap.Int64("author_id", post.CommunityID),
				zap.Error(err))
			continue //这里注意是循环，一个出错直接忽略
		}
		//根据社区id获得社区信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
			continue //这里注意是循环，一个出错直接忽略
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

func GetPostList2(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	//1.去redis查询列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		return
	}
	//2.根据ids去mysql查询帖子详情,返回的顺序还要按照我给定的id的顺序
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	//提前查好每篇帖子的投票数
	voteDate, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}
	//将帖子的作者以及社区等信息填充到帖子详情中
	for idx, post := range posts {
		//根据作者id获取作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById failed",
				zap.Int64("author_id", post.CommunityID),
				zap.Error(err))
			continue //这里注意是循环，一个出错直接忽略
		}
		//根据社区id获得社区信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
			continue //这里注意是循环，一个出错直接忽略
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			VoteNum:         voteDate[idx],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}
func GetCommunityPostList(p *models.ParamCommunityPostList) (data []*models.ApiPostDetail, err error) {
	//1.去redis查询列表
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		return
	}
	//2.根据ids去mysql查询帖子详情,返回的顺序还要按照我给定的id的顺序
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	//提前查好每篇帖子的投票数
	voteDate, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}
	//将帖子的作者以及社区等信息填充到帖子详情中
	for idx, post := range posts {
		//根据作者id获取作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById failed",
				zap.Int64("author_id", post.CommunityID),
				zap.Error(err))
			continue //这里注意是循环，一个出错直接忽略
		}
		//根据社区id获得社区信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
			continue //这里注意是循环，一个出错直接忽略
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			VoteNum:         voteDate[idx],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}
