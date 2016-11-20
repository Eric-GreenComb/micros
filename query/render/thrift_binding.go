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

func (tb thriftBinding) RenderTpl(tplname string, keyMap map[string]string) (string, error) {
	v := tb.RenderService.RenderTpl(tplname, keyMap)
	return v, nil
}
