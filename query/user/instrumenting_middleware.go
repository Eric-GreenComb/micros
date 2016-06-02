package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/query/user/service"
)

type instrumentingMiddleware struct {
	service.UserService
	requestDuration metrics.TimeHistogram
}

func (m instrumentingMiddleware) Ping() (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "Ping"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.UserService.Ping()
	return
}

func (m instrumentingMiddleware) GetUserByEmail(email string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "GetUserByEmail"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.UserService.GetUserByEmail(email)
	return
}

func (m instrumentingMiddleware) GetUserByID(id string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "GetUserByID"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.UserService.GetUserByID(id)
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
