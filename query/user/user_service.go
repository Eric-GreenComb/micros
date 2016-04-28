package main

import (
	"labix.org/v2/mgo/bson"

	"github.com/banerwai/micros/query/user/service"
)

type inmemService struct {
}

func newInmemService() service.UserService {
	return &inmemService{}
}

func (self *inmemService) Ping() string {
	return "pong"
}

func (self *inmemService) GetUser(email string) (r string) {
	var _bson_m bson.M
	err := UsersCollection.Find(bson.M{"email": email}).One(&_bson_m)

	if err != nil {
		return ""
	}

	_data, _err := bson.Marshal(_bson_m)
	if _err != nil {
		return ""
	}

	r = string(_data)
	return
}

func (self *inmemService) CountUser() int64 {
	return 100
}
