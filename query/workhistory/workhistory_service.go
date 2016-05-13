package main

import (
	"github.com/banerwai/micros/query/workhistory/service"
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
	r = profile_id
	return
}
