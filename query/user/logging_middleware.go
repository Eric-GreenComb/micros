package main

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/query/user/service"
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

func (m loggingMiddleware) GetUser(email string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetUser",
			"email", email,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.UserService.GetUser(email)
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
