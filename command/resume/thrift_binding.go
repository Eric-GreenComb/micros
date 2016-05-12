package main

import (
	"github.com/banerwai/micros/command/resume/service"
)

type thriftBinding struct {
	service.ResumeService
}

func (tb thriftBinding) Ping() (string, error) {
	r := tb.ResumeService.Ping()
	return r, nil
}

func (tb thriftBinding) AddResume(json_resume string) (string, error) {
	r := tb.ResumeService.AddResume(json_resume)
	return r, nil
}

func (tb thriftBinding) UpdateResume(json_resume string) (string, error) {
	r := tb.ResumeService.UpdateResume(json_resume)
	return r, nil
}
