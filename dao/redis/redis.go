package redis

// Redis
import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// 声明一个全局的redis 变量
var rdb *redis.Client

// 初始化redis

func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.poolsize"),
	})
	// 初始化redis

	_, err = rdb.Ping().Result()
	if err != nil {
		return
	}
	zap.L().Info("redis init success")
	return
}

func Close() {
	_ = rdb.Close()
}
