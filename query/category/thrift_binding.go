package main

import (
	"github.com/banerwai/micros/query/category/service"
)

type thriftBinding struct {
	service.CategoryService
}

func (tb thriftBinding) Ping() (string, error) {
	return tb.CategoryService.Ping(), nil
}

func (tb thriftBinding) LoadCategory(path string) (bool, error) {
	return tb.CategoryService.LoadCategory(path), nil
}

func (tb thriftBinding) GetCategories() (string, error) {
	return tb.CategoryService.GetCategories(), nil
}

func (tb thriftBinding) GetSubCategories(serialnumber int32) (string, error) {
	return tb.CategoryService.GetSubCategories(serialnumber), nil
}
