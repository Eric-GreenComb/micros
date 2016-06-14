package main

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/query/contact/service"
)

type loggingMiddleware struct {
	service.ContactService
	log.Logger
}

func (m loggingMiddleware) Ping() (v string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "Ping",
			"v", v,
			"took", time.Since(begin),
		)
	}(time.Now())
	v = m.ContactService.Ping()
	return
}

func (m loggingMiddleware) GetContactTpl(tplname string) (v string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetContactTpl",
			"tplname", tplname,
			"v", v,
			"took", time.Since(begin),
		)
	}(time.Now())
	v = m.ContactService.GetContactTpl(tplname)
	return
}

func (m loggingMiddleware) GetContact(contactid string) (v string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetContact",
			"contactid", contactid,
			"v", v,
			"took", time.Since(begin),
		)
	}(time.Now())
	v = m.ContactService.GetContact(contactid)
	return
}

func (m loggingMiddleware) GetContactSignStatus(contactid string) (v string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetContactSignStatus",
			"contactid", contactid,
			"v", v,
			"took", time.Since(begin),
		)
	}(time.Now())
	v = m.ContactService.GetContactSignStatus(contactid)
	return
}

func (m loggingMiddleware) GetClientContact(clientemail string) (v string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetClientContact",
			"clientemail", clientemail,
			"v", v,
			"took", time.Since(begin),
		)
	}(time.Now())
	v = m.ContactService.GetClientContact(clientemail)
	return
}

func (m loggingMiddleware) GetFreelancerContact(freelanceremail string) (v string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "GetFreelancerContact",
			"freelanceremail", freelanceremail,
			"v", v,
			"took", time.Since(begin),
		)
	}(time.Now())
	v = m.ContactService.GetFreelancerContact(freelanceremail)
	return
}
