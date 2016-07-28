package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/query/account/service"
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

func (m instrumentingMiddleware) GetAccountByUserId(user_id string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "GetAccountByUserId"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.AccountService.GetAccountByUserId(user_id)
	return
}

func (m instrumentingMiddleware) GetBillingById(id string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "GetBillingById"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.AccountService.GetBillingById(id)
	return
}

func (m instrumentingMiddleware) GetDealBillingByUserId(user_id string, timestamp int64, pagesize int64) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "GetDealBillingByUserId"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.AccountService.GetDealBillingByUserId(user_id, timestamp, pagesize)
	return
}

func (m instrumentingMiddleware) GetBillingByUserId(user_id string, timestamp int64, pagesize int64) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "GetBillingByUserId"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.AccountService.GetBillingByUserId(user_id, timestamp, pagesize)
	return
}
