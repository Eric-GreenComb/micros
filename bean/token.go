package bean

import (
	"labix.org/v2/mgo/bson"
	"time"
)

// 随机token

// token type
const (
	TokenPwd = iota
	TokenActiveEmail
	TokenUpdateEmail
)

// 过期时间
const (
	PwdOverHours         = 2.0
	ActiveEmailOverHours = 48.0
	UpdateEmailOverHours = 2.0
)

type Token struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Key         string        `bson:"key" json:"key"`
	Token       string        `bson:"token" json:"token"`
	Type        int64         `bson:"type" json:"type"`
	CreatedTime time.Time     `bson:"createdtime" json:"cteatetime"`
}
