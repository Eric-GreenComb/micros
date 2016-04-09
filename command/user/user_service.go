package main

import (
	"github.com/banerwai/micros/command/user/service"

	"github.com/banerwai/gather/command/bean"
	"github.com/banerwai/gommon/db/mongo"
)

type inmemService struct {
}

func newInmemService() service.UserService {
	return &inmemService{}
}

func (self *inmemService) CreateUser(email string, usernameraw string, pwd string) (r string) {
	var _user bean.User
	_user.Email = email
	_user.UsernameRaw = usernameraw
	_user.Pwd = pwd
	mongo.Insert(UsersCollection, _user)
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
