package thrift

import (
	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/profile/service"
	thriftprofile "github.com/banerwai/micros/profile/thrift/gen-go/profile"
)

// New returns an AddService that's backed by the Thrift client.
func New(cli *thriftprofile.ProfileServiceClient, logger log.Logger) service.ProfileService {
	return &client{cli, logger}
}

type client struct {
	*thriftprofile.ProfileServiceClient
	log.Logger
}

func (c client) GetProfile(profile_id string) string {
	reply, err := c.ProfileServiceClient.GetProfile(profile_id)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) GetProfileByCat(name string) string {
	reply, err := c.ProfileServiceClient.GetProfileByCat(name)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) GetProfileBySubCat(name string) string {
	reply, err := c.ProfileServiceClient.GetProfileBySubCat(name)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}
