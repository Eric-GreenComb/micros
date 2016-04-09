package main

import (
	"github.com/banerwai/micros/command/auth/service"
)

type thriftBinding struct {
	service.AuthService
}

func (tb thriftBinding) Register(email string, pwd string, fromUserId string) (string, error) {
	r := tb.AuthService.Register(email, pwd, fromUserId)
	return r, nil
}

func (tb thriftBinding) Login(emailOrUsername string, pwd string) (string, error) {
	r := tb.AuthService.Login(emailOrUsername, pwd)
	return r, nil
}
