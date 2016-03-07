package main

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/email/service"
	thriftemail "github.com/banerwai/micros/email/thrift/gen-go/email"
)

type loggingMiddleware struct {
	service.EmailService
	log.Logger
}

func (m loggingMiddleware) SendEmail(email *thriftemail.Email) (v bool) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "SendEmail",
			"emailto", email.To,
			"subject", email.Subject,
			"v", v,
			"took", time.Since(begin),
		)
	}(time.Now())
	v = m.EmailService.SendEmail(email)
	return
}
