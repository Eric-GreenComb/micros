package main

import (
	"github.com/banerwai/micros/query/render/service"
)

type thriftBinding struct {
	service.RenderService
}

func (tb thriftBinding) Ping() (string, error) {
	v := tb.RenderService.Ping()
	return v, nil
}

func (tb thriftBinding) RenderHello(tmpl, name string) (string, error) {
	v := tb.RenderService.RenderHello(tmpl, name)
	return v, nil
}
