package thrift

import (
	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/user/service"
	thriftuser "github.com/banerwai/micros/user/thrift/gen-go/user"
)

// New returns an AddService that's backed by the Thrift client.
func New(cli *thriftuser.UserServiceClient, logger log.Logger) service.UserService {
	return &client{cli, logger}
}

type client struct {
	*thriftuser.UserServiceClient
	log.Logger
}

func (c client) CreateUser(email string, usernameraw string, pwd string) string {
	reply, err := c.UserServiceClient.CreateUser(email, usernameraw, pwd)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) UpdatePwd(email string, oldpwd string, newpwd string) bool {
	reply, err := c.UserServiceClient.UpdatePwd(email, oldpwd, newpwd)
	if err != nil {
		c.Logger.Log("err", err)
		return false
	}
	return reply
}

func (c client) ActiveUser(token string) bool {
	reply, err := c.UserServiceClient.ActiveUser(token)
	if err != nil {
		c.Logger.Log("err", err)
		return false
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
