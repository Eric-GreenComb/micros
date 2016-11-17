package main

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/query/category/service"
)

type loggingMiddleware struct {
	service.CategoryService
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
	v = m.CategoryService.Ping()
	return
}

func (m loggingMiddleware) LoadCategory(path string) (v bool) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "LoadCategory",
			"path", path,
			"v", v,
			"took", time.Since(begin),
		)
	}(time.Now())
	v = m.CategoryService.LoadCategory(path)
	return
}

func (m loggingMiddleware) GetCategories() (v string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetCategories",
			"took", time.Since(begin),
		)
	}(time.Now())
	v = m.CategoryService.GetCategories()
	return
}

func (m loggingMiddleware) GetSubCategories(serialnumber int32) (v string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetSubCategories",
			"serialnumber", serialnumber,
			"v", fmt.Sprintf("%v", v),
			"took", time.Since(begin),
		)
	}(time.Now())
	v = m.CategoryService.GetSubCategories(serialnumber)
	return
}
