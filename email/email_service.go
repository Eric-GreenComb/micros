package main

import (
	"github.com/banerwai/gommon/net/smtp"
	"github.com/banerwai/micros/email/service"
	thriftemail "github.com/banerwai/micros/email/thrift/gen-go/email"
)

type inmemService struct {
}

func newInmemService() service.EmailService {
	return &inmemService{}
}

func (self *inmemService) SendEmail(email *thriftemail.Email) bool {
	var _email smtp.Email
	_email.Server(email.Host, email.User, email.Password)
	go _email.Send(email.To, email.Subject, email.Body, email.Mailtype)

	return true
}
