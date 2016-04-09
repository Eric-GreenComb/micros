package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/command/user/service"
)

type instrumentingMiddleware struct {
	service.UserService
	requestDuration metrics.TimeHistogram
}

func (m instrumentingMiddleware) CreateUser(email string, usernameraw string, pwd string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "CreateUser"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.UserService.CreateUser(email, usernameraw, pwd)
	return
}

func (m instrumentingMiddleware) UpdatePwd(email string, oldpwd string, newpwd string) (r bool) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "UpdatePwd"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.UserService.UpdatePwd(email, oldpwd, newpwd)
	return
}

func (m instrumentingMiddleware) ActiveUser(token string) (r bool) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "ActiveUser"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.UserService.ActiveUser(token)
	return
}

func (m instrumentingMiddleware) CountUser() (r int64) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "CountUser"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.UserService.CountUser()
	return
}
