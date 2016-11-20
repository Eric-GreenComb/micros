package main

import (
	"time"

	"github.com/banerwai/micros/query/token/service"
	"gopkg.in/mgo.v2/bson"
)

type inmemService struct {
}

func newInmemService() service.TokenService {
	return &inmemService{}
}

func (ims *inmemService) Ping() string {
	return "pong"
}

// return -1 不存在
// return -2 过期
// return -3 db error
// return 1 验证pass

func (ims *inmemService) VerifyToken(key string, ttype int64, overhour float64) int64 {

	var _bsonM bson.M
	err := TokenCollection.Find(bson.M{"key": key, "type": ttype}).One(&_bsonM)

	if err != nil {
		return -1
	}

	// 验证是否过时
	_token := _bsonM["createdtime"]
	_duration := time.Now().Sub(_token.(time.Time))
	if _duration.Hours() > overhour {
		return -2
	}

	return 1
}
