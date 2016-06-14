package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/query/contact/service"
)

type instrumentingMiddleware struct {
	service.ContactService
	requestDuration metrics.TimeHistogram
}

func (m instrumentingMiddleware) Ping() (v string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "Ping"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	v = m.ContactService.Ping()
	return
}

func (m instrumentingMiddleware) GetContactTpl(tplname string) (v string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "GetContactTpl"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	v = m.ContactService.GetContactTpl(tplname)
	return
}

func (m instrumentingMiddleware) GetContact(contactid string) (v string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "GetContact"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	v = m.ContactService.GetContact(contactid)
	return
}

func (m instrumentingMiddleware) GetContactSignStatus(contactid string) (v string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "GetContactSignStatus"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	v = m.ContactService.GetContactSignStatus(contactid)
	return
}

func (m instrumentingMiddleware) GetClientContact(clientemail string) (v string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "GetClientContact"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	v = m.ContactService.GetClientContact(clientemail)
	return
}

func (m instrumentingMiddleware) GetFreelancerContact(freelanceremail string) (v string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "GetFreelancerContact"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	v = m.ContactService.GetFreelancerContact(freelanceremail)
	return
}
