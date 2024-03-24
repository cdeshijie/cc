package logic

import (
	"gin_demo2/dao/redis"
	"gin_demo2/models"
	"go.uber.org/zap"
	"strconv"
)

func VoteForPost(userID int64, p *models.ParamVoteData) error {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	//因为用到Redis了，所以这里放入redis层
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
