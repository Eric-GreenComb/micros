package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/email/service"
	thriftemail "github.com/banerwai/micros/email/thrift/gen-go/email"
)

type instrumentingMiddleware struct {
	service.EmailService
	requestDuration metrics.TimeHistogram
}

func (m instrumentingMiddleware) SendEmail(email *thriftemail.Email) (v bool) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "SendEmail"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	v = m.EmailService.SendEmail(email)
	return
}
