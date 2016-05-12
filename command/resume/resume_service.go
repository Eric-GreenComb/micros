package main

import (
	"github.com/banerwai/micros/command/resume/service"
)

type inmemService struct {
}

func newInmemService() service.ResumeService {
	return &inmemService{}
}

func (self *inmemService) Ping() (r string) {
	r = "pong"
	return
}

func (self *inmemService) AddResume(json_resume string) (r string) {
	r = json_resume
	return
}

func (self *inmemService) UpdateResume(json_resume string) (r string) {
	r = json_resume
	return
}
