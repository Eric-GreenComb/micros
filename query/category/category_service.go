package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"sync"

	"github.com/banerwai/micros/query/category/service"
)

var (
	errInconsistentIDs = errors.New("inconsistent IDs")
	errAlreadyExists   = errors.New("already exists")
	errNotFound        = errors.New("not found")
)

type inmemService struct {
	mtx     sync.RWMutex
	m       map[int]Category
	sortkey []int
}

func newInmemService() service.CategoryService {
	return &inmemService{
		m:       map[int]Category{},
		sortkey: make([]int, 0),
	}
}

func (self *inmemService) Ping() string { return "pong" }

func (self *inmemService) SayHi(name string) string { return "hi," + name }

func (self *inmemService) GetDemoSubCategory(id string) string {

	b, err := json.Marshal(SubCategory{"001", 10, "name-001", "desc-001"})
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (self *inmemService) GetDemoSubCategories(category_id string) string {
	var subs []SubCategory

	subs = append(subs, SubCategory{"001", 1001, "name-001", "desc-001"})
	subs = append(subs, SubCategory{"002", 1002, "name-002", "desc-0012"})

	b, err := json.Marshal(subs)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (self *inmemService) LoadCategory(path string) bool {
	_f, _err := ioutil.ReadFile(path)
	if _err != nil {
		fmt.Println("error:", _err)
		return false
	}
	var categories []Category
	_err = json.Unmarshal(_f, &categories)
	if _err != nil {
		fmt.Println("error:", _err)
		return false
	}

	return self.initCategories(categories)
}

func (self *inmemService) initCategories(categories []Category) bool {
	self.mtx.Lock()
	defer self.mtx.Unlock()

	for k := range self.m {
		delete(self.m, k)
	}

	for _, _category := range categories {
		if _, ok := self.m[int(_category.SerialNumber)]; ok {
			continue // don't overwrite
		}
		self.m[int(_category.SerialNumber)] = _category
	}

	self.sortkey = make([]int, 0)
	for _k, _ := range self.m {
		self.sortkey = append(self.sortkey, _k)
	}
	sort.Ints(self.sortkey)

	return true
}

func (self *inmemService) GetCategories() string {
	self.mtx.RLock()
	defer self.mtx.RUnlock()

	var _categories []Category
	for _, _k := range self.sortkey {
		_cat := self.m[_k]
		_categories = append(_categories, _cat)
	}

	b, err := json.Marshal(_categories)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (self *inmemService) GetSubCategories(serialnumber int32) string {
	self.mtx.RLock()
	defer self.mtx.RUnlock()

	p, ok := self.m[int(serialnumber)]
	if !ok {
		return ""
	}

	b, err := json.Marshal(p.Subcategories)
	if err != nil {
		return err.Error()
	}
	return string(b)
}
