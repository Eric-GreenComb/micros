package main

import (
	"labix.org/v2/mgo/bson"

	"github.com/banerwai/micros/query/resume/service"
)

type inmemService struct {
}

func newInmemService() service.ResumeService {
	return &inmemService{}
}

func (self *inmemService) Ping() (r string) {
	r = "pong"
	return
}

func (self *inmemService) GetResume(userid string) (r string) {
	if !bson.IsObjectIdHex(userid) {
		return ""
	}
	var _bson_m bson.M
	err := ResumeCollection.Find(bson.M{"userid": bson.ObjectIdHex(userid)}).One(&_bson_m)

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
