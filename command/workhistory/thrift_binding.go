package main

import (
	"github.com/banerwai/micros/command/workhistory/service"
)

type thriftBinding struct {
	service.WorkHistoryService
}

func (tb thriftBinding) Ping() (string, error) {
	r := tb.WorkHistoryService.Ping()
	return r, nil
}

func (tb thriftBinding) AddWorkHistory(json_workhistory string) (string, error) {
	r := tb.WorkHistoryService.AddWorkHistory(json_workhistory)
	return r, nil
}

func (tb thriftBinding) UpdateWorkHistory(json_workhistory string) (string, error) {
	r := tb.WorkHistoryService.UpdateWorkHistory(json_workhistory)
	return r, nil
}
