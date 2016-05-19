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

func (tb thriftBinding) UpdateWorkHistory(profile_id, json_workhistory string) (string, error) {
	r := tb.WorkHistoryService.UpdateWorkHistory(profile_id, json_workhistory)
	return r, nil
}
