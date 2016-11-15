package main

import (
	"github.com/banerwai/micros/command/token/service"
)

type thriftBinding struct {
	service.TokenService
}

func (tb thriftBinding) Ping() (string, error) {
	return tb.TokenService.Ping(), nil
}

func (tb thriftBinding) CreateToken(key string, ttype int64) (string, error) {
	return tb.TokenService.CreateToken(key, ttype), nil
}

func (tb thriftBinding) DeleteToken(key string, ttype int64) (bool, error) {
	return tb.TokenService.DeleteToken(key, ttype), nil
}
