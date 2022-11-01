package hashstore

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"shark/model"
	"strconv"
)

func ReadRedis(conn redis.Conn, userId int64) *model.User {
	var key = fmt.Sprintf("user:%d", userId)
	result, _ := redis.StringMap(conn.Do("hgetall", key))
	var uid, _ = strconv.ParseInt(result["uid"], 10, 64)

	var user = model.User{
		Uid:     uid,
		Nick:    result["nick"],
		Address: result["address"],
		Intro:   result["intro"],
	}
	return &user
}

func WriteRedis(conn redis.Conn, user *model.User) {
	if user == nil {
		return
	}

	var (
		key    = fmt.Sprintf("user:%d", user.Uid)
		values = []interface{}{
			key,
		}
	)
	values = append(values, "uid", user.Uid)
	values = append(values, "nick", user.Nick)
	values = append(values, "address", user.Address)
	values = append(values, "intro", user.Intro)
	_, _ = conn.Do("hmset", values...)
}
