package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"sync"

	"github.com/banerwai/micros/query/category/service"
	thriftcategory "github.com/banerwai/micros/query/category/thrift/gen-go/category"
)

var (
	errInconsistentIDs = errors.New("inconsistent IDs")
	errAlreadyExists   = errors.New("already exists")
	errNotFound        = errors.New("not found")
)

type inmemService struct {
	mtx     sync.RWMutex
	m       map[int]thriftcategory.Category
	sortkey []int
}

func newInmemService() service.CategoryService {
	return &inmemService{
		m:       map[int]thriftcategory.Category{},
		sortkey: make([]int, 0),
	}
}

func (self *inmemService) SayHi(name string) string { return "hi," + name }

func (self *inmemService) GetDemoSubCategory(id string) thriftcategory.SubCategory {
	return thriftcategory.SubCategory{"001", 10, "name-001", "desc-001"}
}

func (self *inmemService) GetDemoSubCategories(category_id string) []thriftcategory.SubCategory {
	var subs []thriftcategory.SubCategory

	subs = append(subs, thriftcategory.SubCategory{"001", 1001, "name-001", "desc-001"})
	subs = append(subs, thriftcategory.SubCategory{"002", 1002, "name-002", "desc-0012"})

	return subs
}

func (self *inmemService) LoadCategory(path string) bool {
	_f, _err := ioutil.ReadFile(path)
	if _err != nil {
		fmt.Println("error:", _err)
		return false
	}
	var categories []thriftcategory.Category
	_err = json.Unmarshal(_f, &categories)
	if _err != nil {
		fmt.Println("error:", _err)
		return false
	}

	return self.initCategories(categories)
}

func (self *inmemService) initCategories(categories []thriftcategory.Category) bool {
	self.mtx.Lock()
	defer self.mtx.Unlock()

	for k := range self.m {
		delete(self.m, k)
	}

	for _, _category := range categories {
		if _, ok := self.m[int(_category.Serialnumber)]; ok {
			continue // don't overwrite
		}
		self.m[int(_category.Serialnumber)] = _category
	}

	self.sortkey = make([]int, 0)
	for _k, _ := range self.m {
		self.sortkey = append(self.sortkey, _k)
	}
	sort.Ints(self.sortkey)

	return true
}

func (self *inmemService) GetCategories() []*thriftcategory.Category {
	self.mtx.RLock()
	defer self.mtx.RUnlock()

	var _categories []*thriftcategory.Category
	for _, _k := range self.sortkey {
		_cat := self.m[_k]
		_categories = append(_categories, &_cat)
	}
	return _categories
}

func (self *inmemService) GetSubCategories(serialnumber int32) []*thriftcategory.SubCategory {
	self.mtx.RLock()
	defer self.mtx.RUnlock()

	p, ok := self.m[int(serialnumber)]
	if !ok {
		return nil
	}
	return p.Subcategories
}