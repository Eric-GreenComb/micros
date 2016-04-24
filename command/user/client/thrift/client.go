package thrift

import (
	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/command/user/service"
	thriftuser "github.com/banerwai/micros/command/user/thrift/gen-go/user"
)

// New returns an AddService that's backed by the Thrift client.
func New(cli *thriftuser.UserServiceClient, logger log.Logger) service.UserService {
	return &client{cli, logger}
}

type client struct {
	*thriftuser.UserServiceClient
	log.Logger
}

func (c client) CreateUser(mmap map[string]string) string {
	reply, err := c.UserServiceClient.CreateUser(mmap)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) ResetPwd(email string, newpwd string) bool {
	reply, err := c.UserServiceClient.ResetPwd(email, newpwd)
	if err != nil {
		c.Logger.Log("err", err)
		return false
	}
	return reply
}

func (c client) ActiveUser(email string) bool {
	reply, err := c.UserServiceClient.ActiveUser(email)
	if err != nil {
		c.Logger.Log("err", err)
		return false
	}
	return reply
}
