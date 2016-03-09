package main

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/token/service"
)

type loggingMiddleware struct {
	service.TokenService
	log.Logger
}

func (m loggingMiddleware) NewToken_(key string, ttype int64) (v string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "NewToken_",
			"key", key,
			"ttype", ttype,
			"v", v,
			"took", time.Since(begin),
		)
	}(time.Now())
	v = m.TokenService.NewToken_(key, ttype)
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

func (m loggingMiddleware) VerifyToken(key string, ttype int64) (v int64) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "VerifyToken",
			"key", key,
			"ttype", ttype,
			"v", v,
			"took", time.Since(begin),
		)
	}(time.Now())
	v = m.TokenService.VerifyToken(key, ttype)
	return
}
