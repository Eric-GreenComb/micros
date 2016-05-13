package main

import (
	"github.com/banerwai/micros/command/workhistory/service"
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

func (self *inmemService) AddWorkHistory(json_workhistory string) (r string) {
	r = json_workhistory
	return
}

func (self *inmemService) UpdateWorkHistory(json_workhistory string) (r string) {
	r = json_workhistory
	return
}
