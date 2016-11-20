package main

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/banerwai/micros/query/user/service"
)

type inmemService struct {
}

func newInmemService() service.UserService {
	return &inmemService{}
}

func (ims *inmemService) Ping() string {
	return "pong"
}

func (ims *inmemService) GetUserByEmail(email string) (r string) {
	var _bsonM bson.M
	err := UsersCollection.Find(bson.M{"email": email}).One(&_bsonM)

	if err != nil {
		return ""
	}

	_data, _err := bson.Marshal(_bsonM)
	if _err != nil {
		return ""
	}

	r = string(_data)
	return
}

func (ims *inmemService) GetUserByID(ID string) (r string) {
	if !bson.IsObjectIdHex(ID) {
		return ""
	}
	var _bsonM bson.M
	err := UsersCollection.Find(bson.M{"_id": bson.ObjectIdHex(ID)}).One(&_bsonM)

	if err != nil {
		return ""
	}

	_data, _err := bson.Marshal(_bsonM)
	if _err != nil {
		return ""
	}

	r = string(_data)
	return
}

func (ims *inmemService) CountUser() int64 {
	return 100
}
