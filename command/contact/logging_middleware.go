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

func (m loggingMiddleware) CreateContact(json_contact string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "CreateContact",
			"json_contact", json_contact,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ContactService.CreateContact(json_contact)
	return
}

func (m loggingMiddleware) ClientSignContact(contact_id string, status bool) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "ClientSignContact",
			"contact_id", contact_id,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ContactService.ClientSignContact(contact_id, status)
	return
}

func (m loggingMiddleware) FreelancerSignContact(contact_id string, status bool) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "FreelancerSignContact",
			"contact_id", contact_id,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ContactService.FreelancerSignContact(contact_id, status)
	return
}

func (m loggingMiddleware) DealContact(contact_id string, status bool) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "DealContact",
			"contact_id", contact_id,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ContactService.DealContact(contact_id, status)
	return
}

func (m loggingMiddleware) UpdateContact(contact_id string, mmap map[string]string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "UpdateContact",
			"contact_id", contact_id,
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ContactService.UpdateContact(contact_id, mmap)
	return
}
