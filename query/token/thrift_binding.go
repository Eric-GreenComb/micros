package main

import (
	"github.com/banerwai/micros/query/token/service"
)

type thriftBinding struct {
	service.TokenService
}

func (tb thriftBinding) Ping() (string, error) {
	return tb.TokenService.Ping(), nil
}

func (tb thriftBinding) VerifyToken(key string, ttype int64, overhour int64) (int64, error) {
	return tb.TokenService.VerifyToken(key, ttype, overhour), nil
}
