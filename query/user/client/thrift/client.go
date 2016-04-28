package thrift

import (
	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/query/user/service"
	thriftuser "github.com/banerwai/micros/query/user/thrift/gen-go/user"
)

// New returns an AddService that's backed by the Thrift client.
func New(cli *thriftuser.UserServiceClient, logger log.Logger) service.UserService {
	return &client{cli, logger}
}

type client struct {
	*thriftuser.UserServiceClient
	log.Logger
}

func (c client) Ping() string {
	reply, err := c.UserServiceClient.Ping()
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) GetUser(email string) string {
	reply, err := c.UserServiceClient.GetUser(email)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) CountUser() int64 {
	reply, err := c.UserServiceClient.CountUser()
	if err != nil {
		c.Logger.Log("err", err)
		return -1
	}
	return reply
}
