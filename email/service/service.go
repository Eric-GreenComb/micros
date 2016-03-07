package service

import (
	thriftemail "github.com/banerwai/micros/email/thrift/gen-go/email"
)

// EmailService is the abstract representation of this service.
type EmailService interface {
	SendEmail(email *thriftemail.Email) bool
}
