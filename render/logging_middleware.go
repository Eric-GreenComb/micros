package main

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/render/service"
)

type loggingMiddleware struct {
	service.RenderService
	log.Logger
}

func (m loggingMiddleware) RenderHello(tmpl, name string) (v string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "RenderHello",
			"tmpl", tmpl,
			"name", name,
			"v", v,
			"took", time.Since(begin),
		)
	}(time.Now())
	v = m.RenderService.RenderHello(tmpl, name)
	return
}
