package main

import (
	"encoding/json"
	"github.com/banerwai/global/bean"
	"github.com/banerwai/micros/command/workhistory/service"
	"gopkg.in/mgo.v2/bson"
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

func (ims *inmemService) UpdateWorkHistory(profileID, jsonWorkhistory string) (r string) {
	var _workHistory bean.WorkHistory
	err := json.Unmarshal([]byte(jsonWorkhistory), &_workHistory)
	if err != nil {
		return err.Error()
	}
	_workHistory.ID = ""
	_, _err := WorkHistoryCollection.Upsert(bson.M{"profile_id": bson.ObjectIdHex(profileID)}, bson.M{"$set": _workHistory})
	if _err != nil {
		return _err.Error()
	}
	return "OK"
}
