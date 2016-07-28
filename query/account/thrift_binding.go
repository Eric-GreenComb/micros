package main

import (
	"github.com/banerwai/micros/query/account/service"
)

type thriftBinding struct {
	service.AccountService
}

func (tb thriftBinding) Ping() (string, error) {
	r := tb.AccountService.Ping()
	return r, nil
}

func (tb thriftBinding) GetAccountByUserId(user_id string) (string, error) {
	r := tb.AccountService.GetAccountByUserId(user_id)
	return r, nil
}

func (tb thriftBinding) GetBillingById(id string) (string, error) {
	r := tb.AccountService.GetBillingById(id)
	return r, nil
}

func (tb thriftBinding) GetDealBillingByUserId(user_id string, timestamp int64, pagesize int64) (string, error) {
	r := tb.AccountService.GetDealBillingByUserId(user_id, timestamp, pagesize)
	return r, nil
}

func (tb thriftBinding) GetBillingByUserId(user_id string, timestamp int64, pagesize int64) (string, error) {
	r := tb.AccountService.GetBillingByUserId(user_id, timestamp, pagesize)
	return r, nil
}
