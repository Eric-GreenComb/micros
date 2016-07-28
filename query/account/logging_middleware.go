package main

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/query/account/service"
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

func (m loggingMiddleware) GetAccountByUserId(user_id string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetAccountByUserId",
			"user_id", user_id,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.AccountService.GetAccountByUserId(user_id)
	return
}

func (m loggingMiddleware) GetBillingById(id string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetBillingById",
			"id", id,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.AccountService.GetBillingById(id)
	return
}

func (m loggingMiddleware) GetDealBillingByUserId(user_id string, timestamp int64, pagesize int64) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetDealBillingByUserId",
			"user_id", user_id,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.AccountService.GetDealBillingByUserId(user_id, timestamp, pagesize)
	return
}

func (m loggingMiddleware) GetBillingByUserId(user_id string, timestamp int64, pagesize int64) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetBillingByUserId",
			"user_id", user_id,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.AccountService.GetBillingByUserId(user_id, timestamp, pagesize)
	return
}
