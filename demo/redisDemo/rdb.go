package redisDemo

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
}
func RedisDB() {
	opt, err := redis.ParseURL(os.Getenv("redisInfo"))
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(opt)
	ctx := context.Background()
	rdb.Do(ctx, "set", "k1", "v1") //cmd命令
	rdb.Do(ctx, "set", "b1", true) //cmd命令
	rdb.Set(ctx, "k2", time.Now(), 0)
	res, err := rdb.Do(ctx, "get", "k1").Result()
	res1, _ := rdb.Do(ctx, "get", "b1").Bool()
	res2, _ := rdb.Get(ctx, "k2").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key不存在")
		} else {
			fmt.Println("err:", err)
		}
	} else {
		fmt.Println("res:", res.(string))
		fmt.Println("res:", res1)
		fmt.Println("res:", res2)

	}

}
func RedisPipeLine() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
	value := os.Getenv("redisInfo")
	if value == "" {
		fmt.Println("环境变量 MY_ENV_VAR 不存在")
	} else {
		fmt.Println("环境变量 MY_ENV_VAR 的值是:", value)
	}
}
