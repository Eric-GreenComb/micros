package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/query/category/service"
)

type instrumentingMiddleware struct {
	service.CategoryService
	requestDuration metrics.TimeHistogram
}

func (m instrumentingMiddleware) SayHi(name string) (v string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "concat"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	v = m.CategoryService.SayHi(name)
	return
}

func (m instrumentingMiddleware) GetDemoSubCategory(id string) (v string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "GetDemoSubCategory"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	v = m.CategoryService.GetDemoSubCategory(id)
	return
}

func (m instrumentingMiddleware) GetDemoSubCategories(category_id string) (v string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "GetDemoSubCategories"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	v = m.CategoryService.GetDemoSubCategories(category_id)
	return
}

func (m instrumentingMiddleware) LoadCategory(path string) (v bool) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "LoadCategory"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	v = m.CategoryService.LoadCategory(path)
	return
}

func (m instrumentingMiddleware) GetCategories() (v string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "GetCategories"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	v = m.CategoryService.GetCategories()
	return
}

func (m instrumentingMiddleware) GetSubCategories(serialnumber int32) (v string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "GetSubCategories"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	v = m.CategoryService.GetSubCategories(serialnumber)
	return
}
