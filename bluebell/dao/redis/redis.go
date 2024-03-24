package redis

//用来初始化redis连接
import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var rdb *redis.Client
var ctx = context.Background()
var Nil = redis.Nil

// 初始化redis连接
func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"), // 密码
		DB:       viper.GetInt("redis.db"),          // 数据库
		PoolSize: viper.GetInt("redis.poll_size"),   // 连接池大小
	})
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		zap.L().Error("redis ping err", zap.Error(err))
	}
	return
}
func Close() {
	_ = rdb.Close()
}
