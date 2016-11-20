package main

import (
	"bytes"
	"html/template"
	"sync"

	"github.com/banerwai/global/constant"
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

func (ims *inmemService) Ping() string {
	return "pong"
}

func (ims *inmemService) RenderTpl(tplname string, keyMap map[string]string) string {

	_tpl, err := ims.cachedTpl(tplname)
	if err != nil {
		return ""
	}

	tpl, _ := template.New(tplname).Parse(_tpl)

	b := bytes.NewBuffer(make([]byte, 0))

	err = tpl.Execute(b, keyMap)
	if err != nil {
		return ""
	}
	return b.String()
}

func (ims *inmemService) cachedTpl(tplname string) (string, error) {
	ims.mtx.RLock()
	defer ims.mtx.RUnlock()

	v, ok := ims.m[tplname]
	if ok {
		return v, nil
	}

	_tpl, _err := ims.getTplFromEtcd(tplname)
	if _err != nil {
		return "", _err
	}

	ims.m[tplname] = _tpl

	return _tpl, nil
}

func (ims *inmemService) getTplFromEtcd(tplname string) (string, error) {
	_key := constant.EtcdKeyTplWeb + tplname
	_tpl, _err := etcd.GetValue(_key)
	if _err != nil {
		return "", _err
	}
	return _tpl, nil
}
