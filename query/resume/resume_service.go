package main

import (
	"github.com/banerwai/micros/query/resume/service"
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

func (self *inmemService) GetResume(id string) (r string) {
	r = id
	return
}
