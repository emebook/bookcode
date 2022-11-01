package model

type User struct {
	Uid     int64  `json:"uid"`
	Nick    string `json:"nick"`
	Address string `json:"address"`
	Intro   string `json:"intro"`
}
