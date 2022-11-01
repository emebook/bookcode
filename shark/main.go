package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"shark/hashstore"
	"shark/jsonstore"
	"shark/model"
	"shark/zipstore"
	"strings"
	"time"
)

var (
	redisPool *redis.Pool
	redisConn redis.Conn
)

func init() {
	redisPool = &redis.Pool{
		MaxIdle:     1,
		MaxActive:   500,
		IdleTimeout: time.Second * time.Duration(20),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			return c, err
		},
	}
	redisConn = redisPool.Get()
}

func generateUser() *model.User {
	var (
		uid  int64 = 12345678
		user       = model.User{
			Uid:     uid,
			Nick:    "用户昵称",
			Address: "用户地址",
			Intro:   "这是我的个人介绍",
		}
	)
	user.Intro = strings.Repeat(user.Intro, 1000000)
	return &user
}

func main() {
	var (
		user      = generateUser()
		redisUser *model.User
	)

	// 通过使用三种不同的方式, 观察redis内存的变化

	// hash处理方式
	hashstore.WriteRedis(redisConn, user)
	redisUser = hashstore.ReadRedis(redisConn, user.Uid)
	fmt.Printf("%d : name=%s , address=%s , intro.length=%d\n", redisUser.Uid, redisUser.Nick, redisUser.Address, len([]rune(redisUser.Intro)))

	// JSON处理方式
	jsonstore.WriteRedis(redisConn, user)
	redisUser = jsonstore.ReadRedis(redisConn, user.Uid)
	fmt.Printf("%d : name=%s , address=%s , intro.length=%d\n", redisUser.Uid, redisUser.Nick, redisUser.Address, len([]rune(redisUser.Intro)))

	// Zip+Protobuf处理方式
	zipstore.WriteRedis(redisConn, user)
	redisUser = zipstore.ReadRedis(redisConn, user.Uid)
	fmt.Printf("%d : name=%s , address=%s , intro.length=%d", redisUser.Uid, redisUser.Nick, redisUser.Address, len([]rune(redisUser.Intro)))
}
