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

func (m instrumentingMiddleware) RenderTpl(tplname string, key_mmap map[string]string) (v string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "RenderTpl"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	v = m.RenderService.RenderTpl(tplname, key_mmap)
	return
}
