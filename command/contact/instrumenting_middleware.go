package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/command/contact/service"
)

type instrumentingMiddleware struct {
	service.ContactService
	requestDuration metrics.TimeHistogram
}

func (m instrumentingMiddleware) Ping() (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "Ping"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ContactService.Ping()
	return
}

func (m instrumentingMiddleware) CreateContact(json_contact string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "CreateContact"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ContactService.CreateContact(json_contact)
	return
}

func (m instrumentingMiddleware) ClientSignContact(contact_id string, status bool) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "ClientSignContact"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ContactService.ClientSignContact(contact_id, status)
	return
}

func (m instrumentingMiddleware) FreelancerSignContact(contact_id string, status bool) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "FreelancerSignContact"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ContactService.FreelancerSignContact(contact_id, status)
	return
}

func (m instrumentingMiddleware) DealContact(contact_id string, status bool) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "DealContact"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ContactService.DealContact(contact_id, status)
	return
}

func (m instrumentingMiddleware) UpdateContact(contact_id string, mmap map[string]string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "UpdateContact"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ContactService.UpdateContact(contact_id, mmap)
	return
}
