package main

import (
	"github.com/banerwai/gommon/crypto"
	"github.com/banerwai/micros/query/auth/service"
	"labix.org/v2/mgo/bson"
)

type inmemService struct {
}

func newInmemService() service.AuthService {
	return &inmemService{}
}

func (self *inmemService) Ping() (r string) {
	r = "pong"
	return
}

func (self *inmemService) Login(email string, pwd string) (r string) {
	var _bson_m bson.M
	err := UsersCollection.Find(bson.M{"email": email}).One(&_bson_m)

	if err != nil {
		return "error:" + err.Error()
	}
	_pwd := _bson_m["pwd"].(string)

	_is := crypto.CompareHash([]byte(_pwd), pwd)
	if !_is {
		return "error: compare false"
	}

	_data, _err := bson.Marshal(_bson_m)
	if _err != nil {
		return "error: bson.Marshal error"
	}

	r = string(_data)
	return
}
