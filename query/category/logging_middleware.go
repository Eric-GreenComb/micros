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

func (m loggingMiddleware) SayHi(name string) (v string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "Hi",
			"name", name,
			"v", v,
			"took", time.Since(begin),
		)
	}(time.Now())
	v = m.CategoryService.SayHi(name)
	return
}

func (m loggingMiddleware) GetDemoSubCategory(id string) (v string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetDemoSubCategory",
			"id", id,
			"v", fmt.Sprintf("%v", v),
			"took", time.Since(begin),
		)
	}(time.Now())
	v = m.CategoryService.GetDemoSubCategory(id)
	return
}

func (m loggingMiddleware) GetDemoSubCategories(category_id string) (v string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetDemoSubCategories",
			"category_id", category_id,
			"v", fmt.Sprintf("%v", v),
			"took", time.Since(begin),
		)
	}(time.Now())
	v = m.CategoryService.GetDemoSubCategories(category_id)
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
			"v", fmt.Sprintf("%v", v),
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
