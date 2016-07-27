package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/command/account/service"
)

type instrumentingMiddleware struct {
	service.AccountService
	requestDuration metrics.TimeHistogram
}

func (m instrumentingMiddleware) Ping() (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "Ping"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.AccountService.Ping()
	return
}

func (m instrumentingMiddleware) CreateAccount(json_account string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "CreateAccount"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.AccountService.CreateAccount(json_account)
	return
}

func (m instrumentingMiddleware) CreateBilling(json_billing string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "CreateBilling"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.AccountService.CreateBilling(json_billing)
	return
}

func (m instrumentingMiddleware) DealBilling(billing_id string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "DealBilling"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.AccountService.DealBilling(billing_id)
	return
}

func (m instrumentingMiddleware) CancelBilling(billing_id string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "CancelBilling"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.AccountService.CancelBilling(billing_id)
	return
}

func (m instrumentingMiddleware) GenAccount(user_id string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "GenAccount"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.AccountService.GenAccount(user_id)
	return
}
