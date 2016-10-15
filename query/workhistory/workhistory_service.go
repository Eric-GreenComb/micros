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

func (ims *inmemService) Ping() (r string) {
	r = "pong"
	return
}

func (ims *inmemService) GetWorkHistory(profileID string) (r string) {
	var _bsonM bson.M
	err := WorkHistoryCollection.Find(bson.M{"profile_id": bson.ObjectIdHex(profileID)}).One(&_bsonM)

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
