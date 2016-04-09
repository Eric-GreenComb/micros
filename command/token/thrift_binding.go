package main

import (
	"github.com/banerwai/micros/command/token/service"
)

type thriftBinding struct {
	service.TokenService
}

func (tb thriftBinding) NewToken_(key string, ttype int64) (string, error) {
	return tb.TokenService.NewToken_(key, ttype), nil
}

func (tb thriftBinding) DeleteToken(key string, ttype int64) (bool, error) {
	return tb.TokenService.DeleteToken(key, ttype), nil
}

func (tb thriftBinding) VerifyToken(key string, ttype int64) (int64, error) {
	return tb.TokenService.VerifyToken(key, ttype), nil
}