package jsonstore

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"shark/model"
)

func ReadRedis(conn redis.Conn, userId int64) *model.User {
	var key = fmt.Sprintf("user:%d", userId)
	result, _ := redis.String(conn.Do("get", key))

	var user model.User
	_ = json.Unmarshal([]byte(result), &user)
	return &user
}

func WriteRedis(conn redis.Conn, user *model.User) {
	if user == nil {
		return
	}

	var (
		key    = fmt.Sprintf("user:%d", user.Uid)
		buf, _ = json.Marshal(user)
	)
	_, _ = conn.Do("set", key, string(buf))
}
