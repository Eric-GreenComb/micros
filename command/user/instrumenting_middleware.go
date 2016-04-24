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

func (m instrumentingMiddleware) CreateUser(mmap map[string]string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "CreateUser"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.UserService.CreateUser(mmap)
	return
}

func (m instrumentingMiddleware) ResetPwd(email string, newpwd string) (r bool) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "ResetPwd"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.UserService.ResetPwd(email, newpwd)
	return
}

func (m instrumentingMiddleware) ActiveUser(email string) (r bool) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "ActiveUser"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.UserService.ActiveUser(email)
	return
}
