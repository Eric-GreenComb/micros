package main

import (
	"github.com/banerwai/micros/auth/service"
)

type inmemService struct {
}

func newInmemService() service.AuthService {
	return &inmemService{}
}

func (self *inmemService) Register(email string, pwd string, fromUserId string) (r string) {
	r = email + pwd + fromUserId
	return
}

func (self *inmemService) Login(emailOrUsername string, pwd string) (r string) {
	r = emailOrUsername + pwd
	return
}
