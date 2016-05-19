package main

import (
	"github.com/banerwai/micros/query/workhistory/service"
	"labix.org/v2/mgo/bson"
)

type inmemService struct {
}

func newInmemService() service.WorkHistoryService {
	return &inmemService{}
}

func (self *inmemService) Ping() (r string) {
	r = "pong"
	return
}

func (self *inmemService) GetWorkHistory(profile_id string) (r string) {
	var _bson_m bson.M
	err := WorkHistoryCollection.Find(bson.M{"profile_id": bson.ObjectIdHex(profile_id)}).One(&_bson_m)

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
