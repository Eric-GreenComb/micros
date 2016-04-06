package main

import (
	"github.com/banerwai/micros/user/service"
)

type inmemService struct {
}

func newInmemService() service.UserService {
	return &inmemService{}
}

func (self *inmemService) CreateUser(email string, usernameraw string, pwd string) (r string) {
	r = email + usernameraw + pwd
	return
}

func (self *inmemService) UpdatePwd(email string, oldpwd string, newpwd string) (r bool) {
	r = true
	return
}

func (self *inmemService) ActiveUser(token string) (r bool) {
	r = true
	return
}

func (self *inmemService) CountUser() int64 {
	return 100
}
