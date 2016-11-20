package main

import (
	"github.com/banerwai/micros/query/resume/service"
)

type thriftBinding struct {
	service.ResumeService
}

func (tb thriftBinding) Ping() (string, error) {
	r := tb.ResumeService.Ping()
	return r, nil
}

func (tb thriftBinding) GetResume(userID string) (string, error) {
	r := tb.ResumeService.GetResume(userID)
	return r, nil
}
