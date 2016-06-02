package main

import (
	"github.com/banerwai/micros/query/auth/service"
)

type thriftBinding struct {
	service.AuthService
}

func (tb thriftBinding) Ping() (string, error) {
	r := tb.AuthService.Ping()
	return r, nil
}

func (tb thriftBinding) Login(email string, pwd string) (string, error) {
	r := tb.AuthService.Login(email, pwd)
	return r, nil
}
