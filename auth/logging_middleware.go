package main

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/auth/service"
)

type loggingMiddleware struct {
	service.AuthService
	log.Logger
}

func (m loggingMiddleware) Register(email string, pwd string, fromUserId string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "Register",
			"email", email,
			"pwd", pwd,
			"fromUserId", fromUserId,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.AuthService.Register(email, pwd, fromUserId)
	return
}

func (m loggingMiddleware) Login(emailOrUsername string, pwd string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "Login",
			"emailOrUsername", emailOrUsername,
			"pwd", pwd,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.AuthService.Login(emailOrUsername, pwd)
	return
}
