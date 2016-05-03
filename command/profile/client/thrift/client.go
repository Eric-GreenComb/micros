package thrift

import (
	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/command/profile/service"
	thriftprofile "github.com/banerwai/micros/command/profile/thrift/gen-go/profile"
)

// New returns an AddService that's backed by the Thrift client.
func New(cli *thriftprofile.ProfileServiceClient, logger log.Logger) service.ProfileService {
	return &client{cli, logger}
}

type client struct {
	*thriftprofile.ProfileServiceClient
	log.Logger
}

func (c client) Ping() string {
	reply, err := c.ProfileServiceClient.Ping()
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) AddProfile(json_profile string) string {
	reply, err := c.ProfileServiceClient.AddProfile(json_profile)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) UpdateProfile(json_profile string) string {
	reply, err := c.ProfileServiceClient.UpdateProfile(json_profile)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) DeleteProfile(id string) string {
	reply, err := c.ProfileServiceClient.DeleteProfile(id)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}
