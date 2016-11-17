package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"sync"

	"github.com/banerwai/global/bean"
	"github.com/banerwai/global/constant"
	"github.com/banerwai/gommon/etcd"
	"github.com/banerwai/micros/query/category/service"
)

var (
	errInconsistentIDs = errors.New("inconsistent IDs")
	errAlreadyExists   = errors.New("already exists")
	errNotFound        = errors.New("not found")
)

type inmemService struct {
	mtx     sync.RWMutex
	m       map[int]bean.Category
	sortkey []int
}

func newInmemService() service.CategoryService {
	return &inmemService{
		m:       map[int]bean.Category{},
		sortkey: make([]int, 0),
	}
}

func (ims *inmemService) Ping() string { return "pong" }

func (ims *inmemService) LoadCategory(path string) bool {
	_json, _err := ims.getJSONFromEtcd(path)
	if _err != nil {
		fmt.Println("error:", _err)
		return false
	}
	var categories []bean.Category
	_err = json.Unmarshal([]byte(_json), &categories)
	if _err != nil {
		fmt.Println("error:", _err)
		return false
	}

	return ims.initCategories(categories)
}

func (ims *inmemService) getJSONFromEtcd(jsonname string) (string, error) {
	_key := constant.EtcdKeyJSONCategory + jsonname
	_json, _err := etcd.GetValue(_key)
	if _err != nil {
		return "", _err
	}
	return _json, nil
}

func (ims *inmemService) initCategories(categories []bean.Category) bool {
	ims.mtx.Lock()
	defer ims.mtx.Unlock()

	for k := range ims.m {
		delete(ims.m, k)
	}

	for _, _category := range categories {
		if _, ok := ims.m[int(_category.SerialNumber)]; ok {
			continue // don't overwrite
		}
		ims.m[int(_category.SerialNumber)] = _category
	}

	ims.sortkey = make([]int, 0)
	for _k := range ims.m {
		ims.sortkey = append(ims.sortkey, _k)
	}
	sort.Ints(ims.sortkey)

	return true
}

func (ims *inmemService) GetCategories() string {
	ims.mtx.RLock()
	defer ims.mtx.RUnlock()

	var _categories []bean.Category
	for _, _k := range ims.sortkey {
		_cat := ims.m[_k]
		_categories = append(_categories, _cat)
	}

	b, err := json.Marshal(_categories)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (ims *inmemService) GetSubCategories(serialnumber int32) string {
	ims.mtx.RLock()
	defer ims.mtx.RUnlock()

	p, ok := ims.m[int(serialnumber)]
	if !ok {
		return ""
	}

	b, err := json.Marshal(p.Subcategories)
	if err != nil {
		return err.Error()
	}
	return string(b)
}
