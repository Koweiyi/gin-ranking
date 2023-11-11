package cache

import (
	"context"
	"go-ranking/config"

	"github.com/redis/go-redis/v9"
)

var (
	Rdb  *redis.Client
	Rctx context.Context
)

// Go启动就执行该函数
func init() {

	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisAdress,
		Password: config.RedisPassword,
		DB:       0,
	})

	Rctx = context.Background() //	保存

}

func Zscore(id int, score int) redis.Z {
	return redis.Z{Score: float64(score), Member: id}
}