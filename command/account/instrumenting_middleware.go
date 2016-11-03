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

func (m instrumentingMiddleware) CreateAccount(jsonAccount string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "CreateAccount"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.AccountService.CreateAccount(jsonAccount)
	return
}

func (m instrumentingMiddleware) CreateBilling(jsonBilling string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "CreateBilling"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.AccountService.CreateBilling(jsonBilling)
	return
}

func (m instrumentingMiddleware) DealBilling(billingID string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "DealBilling"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.AccountService.DealBilling(billingID)
	return
}

func (m instrumentingMiddleware) CancelBilling(billingID string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "CancelBilling"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.AccountService.CancelBilling(billingID)
	return
}

func (m instrumentingMiddleware) GenAccount(billingID string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "GenAccount"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.AccountService.GenAccount(billingID)
	return
}
