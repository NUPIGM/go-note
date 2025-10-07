package redisDemo_test

import (
	"main/demo/redisDemo"
	"testing"
)

func TestRdb(t *testing.T) {
	redisDemo.RedisDB()
}

func TestPipeline(t *testing.T) {
	redisDemo.RedisPipeLine()
}
