package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/command/token/service"
)

type instrumentingMiddleware struct {
	service.TokenService
	requestDuration metrics.TimeHistogram
}

func (m instrumentingMiddleware) NewToken_(key string, ttype int64) (v string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "NewToken_"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	v = m.TokenService.NewToken_(key, ttype)
	return
}

func (m instrumentingMiddleware) DeleteToken(key string, ttype int64) (v bool) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "DeleteToken"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	v = m.TokenService.DeleteToken(key, ttype)
	return
}
