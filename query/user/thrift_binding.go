package main

import (
	"github.com/banerwai/micros/query/user/service"
)

type thriftBinding struct {
	service.UserService
}

func (tb thriftBinding) Ping() (string, error) {
	r := tb.UserService.Ping()
	return r, nil
}

func (tb thriftBinding) GetUserByEmail(email string) (string, error) {
	r := tb.UserService.GetUserByEmail(email)
	return r, nil
}

func (tb thriftBinding) GetUserByID(id string) (string, error) {
	r := tb.UserService.GetUserByID(id)
	return r, nil
}

func (tb thriftBinding) CountUser() (int64, error) {
	r := tb.UserService.CountUser()
	return r, nil
}
