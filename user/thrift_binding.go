package main

import (
	"github.com/banerwai/micros/user/service"
)

type thriftBinding struct {
	service.UserService
}

func (tb thriftBinding) CreateUser(email string, usernameraw string, pwd string) (string, error) {
	r := tb.UserService.CreateUser(email, usernameraw, pwd)
	return r, nil
}

func (tb thriftBinding) UpdatePwd(email string, oldpwd string, newpwd string) (bool, error) {
	r := tb.UserService.UpdatePwd(email, oldpwd, newpwd)
	return r, nil
}

func (tb thriftBinding) ActiveUser(token string) (bool, error) {
	r := tb.UserService.ActiveUser(token)
	return r, nil
}

func (tb thriftBinding) CountUser() (int64, error) {
	r := tb.UserService.CountUser()
	return r, nil
}
