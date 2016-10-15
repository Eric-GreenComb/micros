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

func (tb thriftBinding) GetWorkHistory(profileID string) (string, error) {
	r := tb.WorkHistoryService.GetWorkHistory(profileID)
	return r, nil
}
