package main

import (
	"github.com/banerwai/micros/query/user/service"
)

type thriftBinding struct {
	service.UserService
}

func (tb thriftBinding) GetUser(email string) (string, error) {
	r := tb.UserService.GetUser(email)
	return r, nil
}

func (tb thriftBinding) CountUser() (int64, error) {
	r := tb.UserService.CountUser()
	return r, nil
}
