package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/banerwai/micros/category/bean"
	"golang.org/x/net/context"
	"io/ioutil"
	"sort"
	"sync"
)

// ProfileService is a simple CRUD interface for user profiles.
type CategoryService interface {
	LoadCategoriesFile(ctx context.Context, file string) error
	GetSubCategories(ctx context.Context, categoryID string) ([]bean.SubCategory, error)
	GetCategories(ctx context.Context) ([]bean.Category, error)
}

var (
	errInconsistentIDs = errors.New("inconsistent IDs")
	errAlreadyExists   = errors.New("already exists")
	errNotFound        = errors.New("not found")
)

type inmemService struct {
	mtx     sync.RWMutex
	m       map[string]bean.Category
	sortkey []string
}

func newInmemService() CategoryService {
	return &inmemService{
		m:       map[string]bean.Category{},
		sortkey: make([]string, 0),
	}
}

func (self *inmemService) LoadCategoriesFile(ctx context.Context, file string) error {
	_f, _ := ioutil.ReadFile("categories.json")
	var categories []bean.Category
	err := json.Unmarshal(_f, &categories)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}

	return self.InitCategories(ctx, categories)
}

func (self *inmemService) InitCategories(ctx context.Context, categories []bean.Category) error {
	self.mtx.Lock()
	defer self.mtx.Unlock()

	for k := range self.m {
		delete(self.m, k)
	}

	for _, _category := range categories {
		if _, ok := self.m[_category.ID]; ok {
			continue // don't overwrite
		}
		self.m[_category.ID] = _category
	}

	self.sortkey = make([]string, 0)
	for _k, _ := range self.m {
		self.sortkey = append(self.sortkey, _k)
	}
	sort.Strings(self.sortkey)

	return nil
}

func (self *inmemService) GetCategories(ctx context.Context) ([]bean.Category, error) {
	self.mtx.RLock()
	defer self.mtx.RUnlock()

	var _categories []bean.Category
	for _, _k := range self.sortkey {
		_categories = append(_categories, self.m[_k])
	}
	return _categories, nil
}

func (self *inmemService) GetSubCategories(ctx context.Context, categoryID string) ([]bean.SubCategory, error) {
	self.mtx.RLock()
	defer self.mtx.RUnlock()
	p, ok := self.m[categoryID]
	if !ok {
		return []bean.SubCategory{}, errNotFound
	}
	return p.SubCategories, nil
}
