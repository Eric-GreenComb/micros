package main

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/command/account/service"
)

type loggingMiddleware struct {
	service.AccountService
	log.Logger
}

func (m loggingMiddleware) Ping() (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "Ping",
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.AccountService.Ping()
	return
}

func (m loggingMiddleware) CreateAccount(json_account string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "CreateAccount",
			"json_account", json_account,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.AccountService.CreateAccount(json_account)
	return
}

func (m loggingMiddleware) CreateBilling(json_billing string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "CreateBilling",
			"json_billing", json_billing,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.AccountService.CreateBilling(json_billing)
	return
}

func (m loggingMiddleware) DealBilling(billing_id string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "DealBilling",
			"billing_id", billing_id,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.AccountService.DealBilling(billing_id)
	return
}

func (m loggingMiddleware) CancelBilling(billing_id string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "CancelBilling",
			"billing_id", billing_id,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.AccountService.CancelBilling(billing_id)
	return
}

func (m loggingMiddleware) GenAccount(user_id string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GenAccount",
			"user_id", user_id,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.AccountService.GenAccount(user_id)
	return
}
