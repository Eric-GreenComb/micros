package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/query/render/service"
)

type instrumentingMiddleware struct {
	service.RenderService
	requestDuration metrics.TimeHistogram
}

func (m instrumentingMiddleware) Ping() (v string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "Ping"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	v = m.RenderService.Ping()
	return
}

func (m instrumentingMiddleware) RenderHello(tmpl, name string) (v string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "RenderHello"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	v = m.RenderService.RenderHello(tmpl, name)
	return
}
