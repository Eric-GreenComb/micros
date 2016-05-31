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

func (tb thriftBinding) RenderTpl(tplname string, key_mmap map[string]string) (string, error) {
	v := tb.RenderService.RenderTpl(tplname, key_mmap)
	return v, nil
}
