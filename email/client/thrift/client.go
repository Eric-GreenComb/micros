package thrift

import (
	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/email/service"
	thriftemail "github.com/banerwai/micros/email/thrift/gen-go/email"
)

// New returns an AddService that's backed by the Thrift client.
func New(cli *thriftemail.EmailServiceClient, logger log.Logger) service.EmailService {
	return &client{cli, logger}
}

type client struct {
	*thriftemail.EmailServiceClient
	log.Logger
}

func (c client) SendEmail(email *thriftemail.Email) bool {
	reply, err := c.EmailServiceClient.SendEmail(email)
	if err != nil {
		c.Logger.Log("err", err)
		return false
	}
	return reply
}
