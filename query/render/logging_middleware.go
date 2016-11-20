package main

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/query/render/service"
)

type loggingMiddleware struct {
	service.RenderService
	log.Logger
}

func (m loggingMiddleware) Ping() (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "Ping",
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.RenderService.Ping()
	return
}

func (m loggingMiddleware) RenderTpl(tplname string, keyMap map[string]string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "RenderTpl",
			"tplname", tplname,
			"keyMap", keyMap,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.RenderService.RenderTpl(tplname, keyMap)
	return
}
