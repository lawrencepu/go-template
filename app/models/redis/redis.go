package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go-template/config"
	"log"
	"time"
)

// Rdb is redis *client
var Rdb *redis.Client

const RdbKye = "redis"

func Init() {
	rdbConfig, err := config.GetRdbConf(RdbKye)
	if err != nil {
		panic(err)
	}
	Rdb = InitRdb(rdbConfig.DB, rdbConfig.PoolSize, rdbConfig.MaxRetries, rdbConfig.IdleTimeout, rdbConfig.Addr, rdbConfig.Pwd)
}

// 初始化redis
func InitRdb(db, poolSize, maxRetries, idleTimeout int, addr, pwd string) *redis.Client {
	return connRdb(db, poolSize, maxRetries, idleTimeout, addr, pwd)
}

// 连接到redis
func connRdb(db, poolSize, maxRetries, idleTimeout int, addr, pwd string) *redis.Client {
	options := redis.Options{
		Addr:        addr,                                     // Redis地址
		DB:          db,                                       // Redis库
		PoolSize:    poolSize,                                 // Redis连接池大小
		MaxRetries:  maxRetries,                               // 最大重试次数
		IdleTimeout: time.Second * time.Duration(idleTimeout), // 空闲链接超时时间
	}
	if pwd != "" {
		options.Password = pwd
	}
	Rdb := redis.NewClient(&options)
	_, err := Rdb.Ping(context.Background()).Result()
	if err == redis.Nil {
		log.Printf("[rdb] Nil reply returned by Rdb when key does not exist.")
	} else if err != nil {
		log.Printf("[rdb] redis fail, err=%s", err)
		panic(err)
	} else {
		log.Printf("[rdb] redis success")
	}
	return Rdb
}
