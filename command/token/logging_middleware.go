package main

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/command/token/service"
)

type loggingMiddleware struct {
	service.TokenService
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
	v = m.TokenService.Ping()
	return
}

func (m loggingMiddleware) CreateToken(key string, ttype int64) (v string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "CreateToken",
			"key", key,
			"ttype", ttype,
			"v", v,
			"took", time.Since(begin),
		)
	}(time.Now())
	v = m.TokenService.CreateToken(key, ttype)
	return
}

func (m loggingMiddleware) DeleteToken(key string, ttype int64) (v bool) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "DeleteToken",
			"key", key,
			"ttype", ttype,
			"v", v,
			"took", time.Since(begin),
		)
	}(time.Now())
	v = m.TokenService.DeleteToken(key, ttype)
	return
}
