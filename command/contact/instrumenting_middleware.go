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

func (m instrumentingMiddleware) CreateContact(jsonContact string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "CreateContact"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ContactService.CreateContact(jsonContact)
	return
}

func (m instrumentingMiddleware) ClientSignContact(contactID string, status bool) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "ClientSignContact"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ContactService.ClientSignContact(contactID, status)
	return
}

func (m instrumentingMiddleware) FreelancerSignContact(contactID string, status bool) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "FreelancerSignContact"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ContactService.FreelancerSignContact(contactID, status)
	return
}

func (m instrumentingMiddleware) DealContact(contactID string, status bool) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "DealContact"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ContactService.DealContact(contactID, status)
	return
}

func (m instrumentingMiddleware) UpdateContact(contactID string, mmap map[string]string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "UpdateContact"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ContactService.UpdateContact(contactID, mmap)
	return
}
