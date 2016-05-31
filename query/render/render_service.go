package main

import (
	"bytes"
	"html/template"
	"sync"

	"github.com/banerwai/global"
	"github.com/banerwai/gommon/etcd"
	"github.com/banerwai/micros/query/render/service"
)

type inmemService struct {
	mtx sync.RWMutex
	m   map[string]string
}

func newInmemService() service.RenderService {
	return &inmemService{
		m: map[string]string{},
	}
}

func (self *inmemService) Ping() string {
	return "pong"
}

func (self *inmemService) RenderTpl(tplname string, key_mmap map[string]string) string {

	_tpl, err := self.cachedTpl(tplname)
	if err != nil {
		return ""
	}

	tpl, _ := template.New(tplname).Parse(_tpl)

	b := bytes.NewBuffer(make([]byte, 0))

	err = tpl.Execute(b, key_mmap)
	if err != nil {
		return ""
	}
	return b.String()
}

func (self *inmemService) cachedTpl(tplname string) (string, error) {
	self.mtx.RLock()
	defer self.mtx.RUnlock()

	v, ok := self.m[tplname]
	if ok {
		return v, nil
	}

	_tpl, _err := self.getTplFromEtcd(tplname)
	if _err != nil {
		return "", _err
	}

	self.m[tplname] = _tpl

	return _tpl, nil
}

func (self *inmemService) getTplFromEtcd(tplname string) (string, error) {
	_key := global.ETCD_KEY_TPL_WEB + tplname
	_tpl, _err := etcd.GetValue(_key)
	if _err != nil {
		return "", _err
	}
	return _tpl, nil
}
