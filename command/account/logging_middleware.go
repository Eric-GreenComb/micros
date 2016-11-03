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

func (m loggingMiddleware) CreateAccount(jsonAccount string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "CreateAccount",
			"json_account", jsonAccount,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.AccountService.CreateAccount(jsonAccount)
	return
}

func (m loggingMiddleware) CreateBilling(jsonBilling string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "CreateBilling",
			"json_billing", jsonBilling,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.AccountService.CreateBilling(jsonBilling)
	return
}

func (m loggingMiddleware) DealBilling(billingID string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "DealBilling",
			"billing_id", billingID,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.AccountService.DealBilling(billingID)
	return
}

func (m loggingMiddleware) CancelBilling(billingID string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "CancelBilling",
			"billing_id", billingID,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.AccountService.CancelBilling(billingID)
	return
}

func (m loggingMiddleware) GenAccount(userID string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GenAccount",
			"user_id", userID,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.AccountService.GenAccount(userID)
	return
}
