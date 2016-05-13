package main

import (
	"github.com/banerwai/micros/query/workhistory/service"
)

type thriftBinding struct {
	service.WorkHistoryService
}

func (tb thriftBinding) Ping() (string, error) {
	r := tb.WorkHistoryService.Ping()
	return r, nil
}

func (tb thriftBinding) GetWorkHistory(profile_id string) (string, error) {
	r := tb.WorkHistoryService.GetWorkHistory(profile_id)
	return r, nil
}
