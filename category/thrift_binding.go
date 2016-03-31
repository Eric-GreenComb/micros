package main

import (
	"github.com/banerwai/micros/category/service"
	thriftcategory "github.com/banerwai/micros/category/thrift/gen-go/category"
)

type thriftBinding struct {
	service.CategoryService
}

func (tb thriftBinding) SayHi(name string) (string, error) {
	v := tb.CategoryService.SayHi(name)
	return v, nil
}

func (tb thriftBinding) GetDemoSubCategory(id string) (*thriftcategory.SubCategory, error) {
	v := tb.CategoryService.GetDemoSubCategory(id)
	sub := thriftcategory.SubCategory{v.ID, v.Serialnumber, v.Name, v.Desc}
	return &sub, nil
}

func (tb thriftBinding) GetDemoSubCategories(category_id string) ([]*thriftcategory.SubCategory, error) {
	_subs := tb.CategoryService.GetDemoSubCategories(category_id)

	var subs []*thriftcategory.SubCategory

	for _index, _ := range _subs {
		subs = append(subs, &_subs[_index])
	}

	return subs, nil
}

func (tb thriftBinding) LoadCategory(path string) (bool, error) {
	v := tb.CategoryService.LoadCategory(path)
	return v, nil
}

func (tb thriftBinding) GetCategories() ([]*thriftcategory.Category, error) {
	return tb.CategoryService.GetCategories(), nil
}

func (tb thriftBinding) GetSubCategories(serialnumber int32) ([]*thriftcategory.SubCategory, error) {
	return tb.CategoryService.GetSubCategories(serialnumber), nil
}
