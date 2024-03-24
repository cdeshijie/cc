package redis

import (
	"gin_demo2/models"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

// getIDsFromKey 从redis中根据排序方式和page，size获取对应的id列表
func getIDsFromKey(key string, page, size int64) ([]string, error) {
	//确认索引查询的起始点
	start := (page - 1) * size
	end := start + size - 1
	//查询,按分数从大到小，并且指定起始点，redis默认从大到小，所以这里取反序
	return rdb.ZRevRange(ctx, key, start, end).Result()
}

// GetPostIDsInOrder 将排序类型和尺寸传给从redis获取id列表的函数
func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	//从redis获取id
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	//确认索引查询的起始点
	return getIDsFromKey(key, p.Page, p.Size)
}

// GetPostVoteData 根据ids查询每个帖子投赞成票的数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	/*data = make([]int64, 0, len(ids))
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPF + id)
		v := rdb.ZCount(ctx, key, "1", "1").Val()
		data = append(data, v)
	}
	return*/
	//使用pipeline一次发送多条命令，减少rtt（往返时延，表示从发送端发送数据开始，到发送端收到来自接收端的确认（接收端收到数据后便立即发送确认），总共经历的时延）
	pipeLine := rdb.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPF + id)
		pipeLine.ZCount(ctx, key, "1", "1")
	}
	cmders, err := pipeLine.Exec(ctx)
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}

// GetCommunityPostIDsInOrder 按社区查询ids
func GetCommunityPostIDsInOrder(p *models.ParamCommunityPostList) ([]string, error) {
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}
	//orderkey指的是按照分数排序还是时间排序
	//使用zinterstore，把分区的帖子set和存储分数的zset取交集，生成一个新的zset
	//针对新的zset，按照之前的逻辑取数据

	//社区key,这里是存储某个社区下所有帖子id的set
	ckey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityID)))

	//key是交集的名字，排序zset名字+社区id
	key := orderKey + strconv.Itoa(int(p.CommunityID))

	//判断之前是否存在key，利用缓存的key减少zinterstore计算次数,没有再创建
	if rdb.Exists(ctx, key).Val() < 1 { //如果交集不存在
		pipeline := rdb.Pipeline()
		pipeline.ZInterStore(ctx, key, &redis.ZStore{
			Keys:      []string{ckey, orderKey},
			Aggregate: "MAX", //取分数最大值
		})
		pipeline.Expire(ctx, key, 60*time.Second) //60s内再次执行就可以不用查询
		_, err := pipeline.Exec(ctx)
		if err != nil {
			return nil, err
		}
	}
	//如果存在直接根据key查询其所有的帖子的id，即ids
	return getIDsFromKey(key, p.Page, p.Size)
}
