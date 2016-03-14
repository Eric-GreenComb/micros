package thrift

import (
	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/auth/service"
	thriftauth "github.com/banerwai/micros/auth/thrift/gen-go/auth"
)

// New returns an AddService that's backed by the Thrift client.
func New(cli *thriftauth.AuthServiceClient, logger log.Logger) service.AuthService {
	return &client{cli, logger}
}

type client struct {
	*thriftauth.AuthServiceClient
	log.Logger
}

func (c client) Register(email string, pwd string, fromUserId string) string {
	reply, err := c.AuthServiceClient.Register(email, pwd, fromUserId)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) Login(emailOrUsername string, pwd string) string {
	reply, err := c.AuthServiceClient.Login(emailOrUsername, pwd)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}
