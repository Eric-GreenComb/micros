package main

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/category/service"
	thriftcategory "github.com/banerwai/micros/category/thrift/gen-go/category"
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

func (m loggingMiddleware) GetDemoSubCategory(id string) (v thriftcategory.SubCategory) {
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

func (m loggingMiddleware) GetDemoSubCategories(category_id string) (v []thriftcategory.SubCategory) {
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

func (m loggingMiddleware) GetCategories() (v []*thriftcategory.Category) {
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

func (m loggingMiddleware) GetSubCategories(category_id string) (v []*thriftcategory.SubCategory) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetSubCategories",
			"category_id", category_id,
			"v", fmt.Sprintf("%v", v),
			"took", time.Since(begin),
		)
	}(time.Now())
	v = m.CategoryService.GetSubCategories(category_id)
	return
}
