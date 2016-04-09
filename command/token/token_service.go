package main

import (
	"fmt"
	"time"

	"github.com/banerwai/gather/command/bean"
	"github.com/banerwai/gommon/uuid"
	"github.com/banerwai/micros/command/token/service"

	"labix.org/v2/mgo/bson"
)

type inmemService struct {
}

func newInmemService() service.TokenService {
	return &inmemService{}
}

func (self *inmemService) NewToken_(key string, ttype int64) string {
	token := bean.Token{Key: key, Token: uuid.UUID(), CreatedTime: time.Now(), Type: ttype}
	_info, _err := mgoCollectionToken.Upsert(bson.M{"key": key, "type": ttype}, token)
	fmt.Println(_info)
	if _err != nil {
		return _err.Error()
	}
	return token.Token
}

func (self *inmemService) DeleteToken(key string, ttype int64) bool {
	mgoCollectionToken.Remove(bson.M{"key": key, "type": ttype})
	return true
}

// return -1 不存在
// return -2 过期
// return 1 验证pass

func (self *inmemService) VerifyToken(key string, ttype int64) int64 {
	_overHours := self.GetOverHours(ttype)
	fmt.Println(_overHours)

	var _token bean.Token

	err := mgoCollectionToken.Find(bson.M{"key": key, "type": ttype}).One(&_token)

	if err != nil {
		return -1
	}

	// 验证是否过时
	_now := time.Now()
	_duration := _now.Sub(_token.CreatedTime)

	if _duration.Hours() > _overHours {
		return -2
	}

	return 1
}

func (self *inmemService) GetOverHours(ttype int64) float64 {
	switch ttype {
	case bean.TokenPwd:
		return bean.PwdOverHours
	case bean.TokenUpdateEmail:
		return bean.UpdateEmailOverHours
	case bean.TokenActiveEmail:
		return bean.ActiveEmailOverHours
	}
	return 0
}
