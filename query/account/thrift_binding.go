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

func (tb thriftBinding) GetAccountByUserID(userID string) (string, error) {
	r := tb.AccountService.GetAccountByUserID(userID)
	return r, nil
}

func (tb thriftBinding) GetBillingByID(ID string) (string, error) {
	r := tb.AccountService.GetBillingByID(ID)
	return r, nil
}

func (tb thriftBinding) GetDealBillingByUserID(userID string, timestamp int64, pagesize int64) (string, error) {
	r := tb.AccountService.GetDealBillingByUserID(userID, timestamp, pagesize)
	return r, nil
}

func (tb thriftBinding) GetBillingByUserID(userID string, timestamp int64, pagesize int64) (string, error) {
	r := tb.AccountService.GetBillingByUserID(userID, timestamp, pagesize)
	return r, nil
}
