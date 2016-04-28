package main

import (
	"time"

	"github.com/banerwai/micros/query/token/service"
	"labix.org/v2/mgo/bson"
)

type inmemService struct {
}

func newInmemService() service.TokenService {
	return &inmemService{}
}

func (self *inmemService) Ping() string {
	return "pong"
}

// return -1 不存在
// return -2 过期
// return -3 db error
// return 1 验证pass

func (self *inmemService) VerifyToken(key string, ttype int64, overhour float64) int64 {

	var _bson_m bson.M
	err := TokenCollection.Find(bson.M{"key": key, "type": ttype}).One(&_bson_m)

	if err != nil {
		return -1
	}

	// 验证是否过时
	_token := _bson_m["createdtime"]
	_duration := time.Now().Sub(_token.(time.Time))
	if _duration.Hours() > overhour {
		return -2
	}

	return 1
}
