package main

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/command/contact/service"
)

type loggingMiddleware struct {
	service.ContactService
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
	r = m.ContactService.Ping()
	return
}

func (m loggingMiddleware) CreateContact(jsonContact string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "CreateContact",
			"jsonContact", jsonContact,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ContactService.CreateContact(jsonContact)
	return
}

func (m loggingMiddleware) ClientSignContact(contactID string, status bool) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "ClientSignContact",
			"contactID", contactID,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ContactService.ClientSignContact(contactID, status)
	return
}

func (m loggingMiddleware) FreelancerSignContact(contactID string, status bool) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "FreelancerSignContact",
			"contactID", contactID,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ContactService.FreelancerSignContact(contactID, status)
	return
}

func (m loggingMiddleware) DealContact(contactID string, status bool) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "DealContact",
			"contactID", contactID,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ContactService.DealContact(contactID, status)
	return
}

func (m loggingMiddleware) UpdateContact(contactID string, mmap map[string]string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "UpdateContact",
			"contactID", contactID,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ContactService.UpdateContact(contactID, mmap)
	return
}
