package main

import (
	"github.com/banerwai/micros/query/auth/service"
)

type thriftBinding struct {
	service.AuthService
}

func (tb thriftBinding) Login(emailOrUsername string, pwd string) (string, error) {
	r := tb.AuthService.Login(emailOrUsername, pwd)
	return r, nil
}
