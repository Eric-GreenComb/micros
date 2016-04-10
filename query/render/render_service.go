package main

import (
	"bytes"
	"fmt"
	"sync"

	"html/template"

	gfile "github.com/banerwai/gommon/file"
	"github.com/banerwai/micros/query/render/service"
)

type inmemService struct {
	mtx sync.RWMutex
	m   map[string]template.Template
}

func newInmemService() service.RenderService {
	return &inmemService{
		m: map[string]template.Template{},
	}
}

type Hello struct {
	Name string
}

func (self *inmemService) RenderHello(tmpl, name string) string {

	t, err := self.cachedTmpl(tmpl)

	if err != nil {
		return ""
	}

	b := bytes.NewBuffer(make([]byte, 0))
	_hello := Hello{Name: name}
	t.Execute(b, _hello)
	return b.String()
}

func (self *inmemService) cachedTmpl(tmpl string) (*template.Template, error) {
	self.mtx.RLock()
	defer self.mtx.RUnlock()

	fmt.Println(tmpl)
	v, ok := self.m[tmpl]
	if ok {
		fmt.Println("cachedTmpl")
		return &v, nil
	}

	_file := gfile.GetCurrentDirectory() + "/tmpl/" + tmpl + ".tmpl"
	fmt.Println(_file)

	t, err := template.ParseFiles(_file)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	self.m[tmpl] = *t

	return t, nil
}
