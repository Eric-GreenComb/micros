package main

import (
	"encoding/json"
	"github.com/banerwai/global/bean"
	"github.com/banerwai/micros/command/workhistory/service"
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

func (self *inmemService) UpdateWorkHistory(profile_id, json_workhistory string) (r string) {
	var _work_history bean.WorkHistory
	err := json.Unmarshal([]byte(json_workhistory), &_work_history)
	if err != nil {
		return err.Error()
	}
	_work_history.Id = ""
	_, _err := WorkHistoryCollection.Upsert(bson.M{"profile_id": bson.ObjectIdHex(profile_id)}, bson.M{"$set": _work_history})
	if _err != nil {
		return _err.Error()
	}
	return "OK"
}
