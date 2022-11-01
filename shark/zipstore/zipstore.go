package zipstore

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/golang/protobuf/proto"
	"io"
	"shark/model"
	"shark/pb"
)

func Encode(src string) string {
	var output bytes.Buffer
	w := zlib.NewWriter(&output)
	_, _ = w.Write([]byte(src))
	_ = w.Close()
	return string(output.Bytes())
}

func Decode(zipContent string) string {
	b := bytes.NewReader([]byte(zipContent))
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	_, _ = io.Copy(&out, r)
	return string(out.Bytes())
}

func ReadRedis(conn redis.Conn, userId int64) *model.User {
	var (
		userPB pb.User
		key    = fmt.Sprintf("user:%d", userId)
	)

	result, _ := redis.String(conn.Do("get", key))
	_ = proto.Unmarshal([]byte(result), &userPB)

	var user = model.User{
		Uid:     *userPB.Uid,
		Nick:    *userPB.Nick,
		Address: *userPB.Address,
		Intro:   string(Decode(*userPB.Intro)),
	}

	return &user
}

func WriteRedis(conn redis.Conn, user *model.User) {
	if user == nil {
		return
	}

	var (
		intro  = Encode(user.Intro)
		userPB = &pb.User{
			Uid:     &user.Uid,
			Nick:    &user.Nick,
			Intro:   &intro,
			Address: &user.Address,
		}
		key    = fmt.Sprintf("user:%d", user.Uid)
		buf, _ = proto.Marshal(userPB)
	)
	_, _ = conn.Do("set", key, string(buf))
}
