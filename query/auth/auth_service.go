package main

import (
	"github.com/banerwai/gommon/crypto"
	"github.com/banerwai/micros/query/auth/service"
	"gopkg.in/mgo.v2/bson"
)

type inmemService struct {
}

func newInmemService() service.AuthService {
	return &inmemService{}
}

func (ims *inmemService) Ping() (r string) {
	r = "pong"
	return
}

func (ims *inmemService) Login(email string, pwd string) (r string) {
	var _bsonM bson.M
	err := UsersCollection.Find(bson.M{"email": email}).One(&_bsonM)

	if err != nil {
		return "error:" + err.Error()
	}

	_active := _bsonM["actived"].(bool)
	if !_active {
		return "error: user email need active"
	}

	_pwd := _bsonM["pwd"].(string)

	_is := crypto.CompareHash([]byte(_pwd), pwd)
	if !_is {
		return "error: compare false"
	}

	_data, _err := bson.Marshal(_bsonM)
	if _err != nil {
		return "error: bson.Marshal error"
	}

	r = string(_data)
	return
}
