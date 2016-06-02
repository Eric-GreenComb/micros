package main

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/query/auth/service"
)

type loggingMiddleware struct {
	service.AuthService
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
	r = m.AuthService.Ping()
	return
}

func (m loggingMiddleware) Login(email string, pwd string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "Login",
			"email", email,
			"pwd", pwd,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.AuthService.Login(email, pwd)
	return
}
