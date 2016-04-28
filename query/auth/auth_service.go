package main

import (
	"github.com/banerwai/micros/query/auth/service"
)

type inmemService struct {
}

func newInmemService() service.AuthService {
	return &inmemService{}
}

func (self *inmemService) Ping() (r string) {
	r = "pong"
	return
}

func (self *inmemService) Login(emailOrUsername string, pwd string) (r string) {
	r = emailOrUsername + pwd
	return
}
