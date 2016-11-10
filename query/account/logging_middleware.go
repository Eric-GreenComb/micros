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

func (m loggingMiddleware) GetAccountByUserID(userID string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetAccountByUserID",
			"userID", userID,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.AccountService.GetAccountByUserID(userID)
	return
}

func (m loggingMiddleware) GetBillingByID(ID string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetBillingByID",
			"ID", ID,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.AccountService.GetBillingByID(ID)
	return
}

func (m loggingMiddleware) GetDealBillingByUserID(userID string, timestamp int64, pagesize int64) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetDealBillingByUserID",
			"userID", userID,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.AccountService.GetDealBillingByUserID(userID, timestamp, pagesize)
	return
}

func (m loggingMiddleware) GetBillingByUserID(userID string, timestamp int64, pagesize int64) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetBillingByUserID",
			"userID", userID,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.AccountService.GetBillingByUserID(userID, timestamp, pagesize)
	return
}
