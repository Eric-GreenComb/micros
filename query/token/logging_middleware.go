package main

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/query/token/service"
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

func (m loggingMiddleware) VerifyToken(key string, ttype int64, overhour int64) (v int64) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "VerifyToken",
			"key", key,
			"ttype", ttype,
			"v", v,
			"took", time.Since(begin),
		)
	}(time.Now())
	v = m.TokenService.VerifyToken(key, ttype, overhour)
	return
}
