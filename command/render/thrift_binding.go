package main

import (
	"github.com/banerwai/micros/command/render/service"
)

type thriftBinding struct {
	service.RenderService
}

func (tb thriftBinding) RenderHello(tmpl, name string) (string, error) {
	v := tb.RenderService.RenderHello(tmpl, name)
	return v, nil
}
