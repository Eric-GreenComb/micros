package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/auth/service"
)

type instrumentingMiddleware struct {
	service.AuthService
	requestDuration metrics.TimeHistogram
}

func (m instrumentingMiddleware) Register(email string, pwd string, fromUserId string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "Register"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.AuthService.Register(email, pwd, fromUserId)
	return
}

func (m instrumentingMiddleware) Login(emailOrUsername string, pwd string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "Login"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.AuthService.Login(emailOrUsername, pwd)
	return
}
