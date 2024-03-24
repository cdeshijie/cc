package redis

import (
	"errors"
	"github.com/redis/go-redis/v9"
	"math"
	"strconv"
	"time"
)

const (
	oneWeekInSeconds = 3600 * 24 * 7
	scorePerVote     = 432 //每一张票的分数
)

var (
	ErrVoteTimeExpire = errors.New("超出投票时间")
	ErrVoteRepeated   = errors.New("投票重复")
)

/*func CreatePost(postID int64) error { //测试不用事务的时候
	//这里向redis进行写入测试
	rdb.Set(ctx, "test2", "cc", time.Hour/2)
	//记录时间，这个是不动的
	fmt.Println("开始在redis中创建帖子", getRedisKey(KeyPostTimeZSet))
	_, err := rdb.ZAdd(ctx, getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	}).Result()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("redis-CreatePost-rdb.ZAdd1成功")
	//记录分数,随着点赞的操作发生变化
	_, err = rdb.ZAdd(ctx, getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	}).Result()
	return err
}*/

func CreatePost(postID, communityID int64) error {
	//这里启用redis的事务，创建帖子和记录分数要同时完成,创建帖子时redis存入两个zset和一个set，一个根据时间的zset，一个根据分数的zset，还有一个set是记录该社区下的所有帖子的id
	pipeline := rdb.TxPipeline()
	//记录时间，这个是不动的
	pipeline.ZAdd(ctx, getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	//记录分数,随着点赞的操作发生变化
	pipeline.ZAdd(ctx, getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	//把帖子id加入到相应set中
	ckey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityID)))
	pipeline.SAdd(ctx, ckey, postID)
	_, err := pipeline.Exec(ctx)
	return err
}

// VoteForPost 为帖子投票的函数
// 本项目使用简化版的投票算法
// 投一张票就加432分 432=86400/200 86400是一天的秒数，这意味着每天需要200张赞成票才能继续靠前
/*
投票的几种情况：
direction=1时：
	1.之前没有投过票，现在投赞成票，分数加432    差值的绝对值1   value > ov  +432
	2.之前投反对票，现在投赞成票,分数加432*2		差值的绝对值2	value > ov	+432*2
direction=0时：
	1.之前投的是赞成票							差值的绝对值1	value < ov	-432
	2.之前投的是反对票							差值的绝对值1	value > ov  +432
direction=2时：
	1.之前没有投过票，现在投反对票  				差值的绝对值1	value < ov	-432
	2.之前投的是赞成票，现在投反对票				差值的绝对值2	value < ov	-432*2
但是不管那种情况，处理流程都是更新分数并且更改投票记录

投票限制：帖子只有在发布的一个星期内才能投票，过期不行
	1.到期之后将redis存储的投票数据放到mysql中
	2.到期之后删除对应的KeyPostVotedZSet
*/

func VoteForPost(userID, postID string, value float64) error {
	//1.判断投票限制,通过KeyPostTimeZSet这个Redis的集合来判断是否过期
	postTime := rdb.ZScore(ctx, getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	//2.更新帖子分数,
	//先查当前用户给当前帖子之前的投票记录，没有投票取出来就是0，这里具体分值计算参上
	ov := rdb.ZScore(ctx, getRedisKey(KeyPostVotedZSetPF+postID), userID).Val() //这里每个帖子都是一个集合
	if value == ov {
		return ErrVoteRepeated
	}
	var op float64
	if value > ov { //如果新值大于旧值，则方向就是正的
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(value - ov)
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(ctx, getRedisKey(KeyPostScoreZSet), diff*op*scorePerVote, postID)
	//3.记录用户对该帖子投票的详情
	if value == 0 { //如果value值为0，说明要取消点赞，这里移除该用户对帖子点赞的记录
		pipeline.ZRem(ctx, getRedisKey(KeyPostVotedZSetPF+postID), userID)
	} else { //其他情况直接覆盖原有记录
		pipeline.ZAdd(ctx, getRedisKey(KeyPostVotedZSetPF+postID), redis.Z{
			Score:  value,
			Member: userID,
		})
	}
	_, err := pipeline.Exec(ctx)
	return err
}
