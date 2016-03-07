package main

import (
	"github.com/banerwai/micros/email/service"
	thriftemail "github.com/banerwai/micros/email/thrift/gen-go/email"
)

type thriftBinding struct {
	service.EmailService
}

func (tb thriftBinding) SendEmail(email *thriftemail.Email) (bool, error) {
	v := tb.EmailService.SendEmail(email)
	return v, nil
}
