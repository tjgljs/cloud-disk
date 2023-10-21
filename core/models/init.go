package models

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func Init(datasource string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", datasource)
	if err != nil {
		fmt.Printf(" xorm init err: %v\n", err)
		return nil
	}
	return engine

}

func InitRedis(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
}
