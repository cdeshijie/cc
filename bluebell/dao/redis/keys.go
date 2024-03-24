package redis

// redis key,这里把key进行拆分，后面可以进行快速的查询和区分
const (
	Prefix             = "bluebell:"
	KeyPostTimeZSet    = "post:time"   //zset，帖子及其发帖时间
	KeyPostScoreZSet   = "post:score"  //zset，帖子及其投票分数
	KeyPostVotedZSetPF = "post:voted:" //zset,记录用户是否投票且投票类型,这里有prefix，说明后面还有参数，这里是post_id
	KeyCommunitySetPF  = "community:"  //保存每个分区下所有的帖子id
)

// 给redis的key加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}
