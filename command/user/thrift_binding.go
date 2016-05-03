package main

import (
	"github.com/banerwai/micros/command/user/service"
)

type thriftBinding struct {
	service.UserService
}

func (tb thriftBinding) Ping() (string, error) {
	r := tb.UserService.Ping()
	return r, nil
}

func (tb thriftBinding) CreateUser(mmap map[string]string) (string, error) {
	r := tb.UserService.CreateUser(mmap)
	return r, nil
}

func (tb thriftBinding) ResetPwd(email string, newpwd string) (bool, error) {
	r := tb.UserService.ResetPwd(email, newpwd)
	return r, nil
}

func (tb thriftBinding) ActiveUser(email string) (bool, error) {
	r := tb.UserService.ActiveUser(email)
	return r, nil
}
