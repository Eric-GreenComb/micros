package service

import (
	thriftcategory "github.com/banerwai/micros/category/thrift/gen-go/category"
)

// CategoryService is the abstract representation of this service.
type CategoryService interface {
	SayHi(name string) string
	GetDemoSubCategory(id string) thriftcategory.SubCategory
	GetDemoSubCategories(category_id string) []thriftcategory.SubCategory

	LoadCategory(path string) bool
	GetCategories() []*thriftcategory.Category
	GetSubCategories(category_id string) []*thriftcategory.SubCategory
}
