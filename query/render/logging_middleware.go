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

func (m loggingMiddleware) Ping() (v string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "Ping",
			"v", v,
			"took", time.Since(begin),
		)
	}(time.Now())
	v = m.RenderService.Ping()
	return
}

func (m loggingMiddleware) RenderTpl(tplname string, key_mmap map[string]string) (v string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "RenderTpl",
			"tplname", tplname,
			"key_mmap", key_mmap,
			"v", v,
			"took", time.Since(begin),
		)
	}(time.Now())
	v = m.RenderService.RenderTpl(tplname, key_mmap)
	return
}
