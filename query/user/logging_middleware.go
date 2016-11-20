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

func (m loggingMiddleware) GetUserByEmail(email string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetUserByEmail",
			"email", email,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.UserService.GetUserByEmail(email)
	return
}

func (m loggingMiddleware) GetUserByID(ID string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetUserByID",
			"ID", ID,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.UserService.GetUserByID(ID)
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
