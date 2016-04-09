package main

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/command/user/service"
)

type loggingMiddleware struct {
	service.UserService
	log.Logger
}

func (m loggingMiddleware) CreateUser(email string, usernameraw string, pwd string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "CreateUser",
			"email", email,
			"usernameraw", usernameraw,
			"pwd", pwd,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.UserService.CreateUser(email, usernameraw, pwd)
	return
}

func (m loggingMiddleware) UpdatePwd(email string, oldpwd string, newpwd string) (r bool) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "UpdatePwd",
			"email", email,
			"oldpwd", oldpwd,
			"newpwd", newpwd,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.UserService.UpdatePwd(email, oldpwd, newpwd)
	return
}

func (m loggingMiddleware) ActiveUser(token string) (r bool) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "ActiveUser",
			"token", token,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.UserService.ActiveUser(token)
	return
}

func (m loggingMiddleware) CountUser() (r int64) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "CountUser",
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.UserService.CountUser()
	return
}
