package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

//redis测试

//声明一个全局的redisdb变量
var redisdb *redis.Client

//初始化连接
func initClient() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               "192.168.42.133:6379",
		Dialer:             nil,
		OnConnect:          nil,
		Password:           "", //no password set
		DB:                 0,  //set default DB
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	})
	_, err = redisdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initClient()
	if err != nil {
		fmt.Printf("client connect to redis failed,err:%v\n", err)
		return
	}
	//设置值
	err = redisdb.Set("ll", "12", 0).Err()
	if err != nil {
		fmt.Printf("redis set key \"ll\" failed,err:%v\n", err)
		return
	}
	//获取值
	xuleVal, err := redisdb.Get("ll").Result()
	if err != nil {
		fmt.Printf("redis get key \"ll\" failed,err:%v\n", err)
		return
	}
	fmt.Printf("ll:val=%v\n", xuleVal)
}
