package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "127.0.0.1:6379",
	Password: "",
	DB:       0,
})

func TestRedisSet(t *testing.T) {
	err := rdb.Set(ctx, "key", "value", time.Second*10).Err()
	if err != nil {
		t.Error(err)
	}
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("连接redis出错，错误信息：%v", err)
	}
	fmt.Println("成功连接redis", pong)
}
func TestRedisGet(t *testing.T) {
	s, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		t.Error(err)
	}
	t.Log(s)
}
