package redis

// Redis
import (
	"fmt"
	"github.com/go-redis/redis"
	"web_app/settings"
)

// 声明一个全局的redis 变量
var rdb *redis.Client

// 初始化redis

func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.Database,
		PoolSize: cfg.PoolSize,
	})
	// 初始化redis

	_, err = rdb.Ping().Result()
	if err != nil {
		return
	}
	return
}

func Close() {
	_ = rdb.Close()
}
