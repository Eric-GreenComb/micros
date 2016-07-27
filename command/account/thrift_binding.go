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

func (tb thriftBinding) CreateAccount(json_account string) (string, error) {
	r := tb.AccountService.CreateAccount(json_account)
	return r, nil
}

func (tb thriftBinding) CreateBilling(json_billing string) (string, error) {
	r := tb.AccountService.CreateBilling(json_billing)
	return r, nil
}

func (tb thriftBinding) DealBilling(billing_id string) (string, error) {
	r := tb.AccountService.DealBilling(billing_id)
	return r, nil
}

func (tb thriftBinding) CancelBilling(billing_id string) (string, error) {
	r := tb.AccountService.CancelBilling(billing_id)
	return r, nil
}

func (tb thriftBinding) GenAccount(user_id string) (string, error) {
	r := tb.AccountService.GenAccount(user_id)
	return r, nil
}
