package main

import (
	"github.com/banerwai/micros/command/account/service"
)

type thriftBinding struct {
	service.AccountService
}

func (tb thriftBinding) Ping() (string, error) {
	r := tb.AccountService.Ping()
	return r, nil
}

func (tb thriftBinding) CreateAccount(jsonAccount string) (string, error) {
	r := tb.AccountService.CreateAccount(jsonAccount)
	return r, nil
}

func (tb thriftBinding) CreateBilling(jsonBilling string) (string, error) {
	r := tb.AccountService.CreateBilling(jsonBilling)
	return r, nil
}

func (tb thriftBinding) DealBilling(billingID string) (string, error) {
	r := tb.AccountService.DealBilling(billingID)
	return r, nil
}

func (tb thriftBinding) CancelBilling(billingID string) (string, error) {
	r := tb.AccountService.CancelBilling(billingID)
	return r, nil
}

func (tb thriftBinding) GenAccount(userID string) (string, error) {
	r := tb.AccountService.GenAccount(userID)
	return r, nil
}
