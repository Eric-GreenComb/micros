package main

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/banerwai/micros/query/resume/service"
)

type inmemService struct {
}

func newInmemService() service.ResumeService {
	return &inmemService{}
}

func (ims *inmemService) Ping() (r string) {
	r = "pong"
	return
}

func (ims *inmemService) GetResume(userID string) (r string) {
	if !bson.IsObjectIdHex(userID) {
		return ""
	}
	var _bsonM bson.M
	err := ResumeCollection.Find(bson.M{"userid": bson.ObjectIdHex(userID)}).One(&_bsonM)

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
