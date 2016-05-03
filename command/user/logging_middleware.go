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

func (m loggingMiddleware) Ping() (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "Ping",
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.UserService.Ping()
	return
}

func (m loggingMiddleware) CreateUser(mmap map[string]string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "CreateUser",
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.UserService.CreateUser(mmap)
	return
}

func (m loggingMiddleware) ResetPwd(email string, newpwd string) (r bool) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "ResetPwd",
			"email", email,
			"newpwd", newpwd,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.UserService.ResetPwd(email, newpwd)
	return
}

func (m loggingMiddleware) ActiveUser(email string) (r bool) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "ActiveUser",
			"email", email,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.UserService.ActiveUser(email)
	return
}
