package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/query/token/service"
)

type instrumentingMiddleware struct {
	service.TokenService
	requestDuration metrics.TimeHistogram
}

func (m instrumentingMiddleware) VerifyToken(key string, ttype int64, overhour float64) (v int64) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "VerifyToken"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	v = m.TokenService.VerifyToken(key, ttype, overhour)
	return
}
