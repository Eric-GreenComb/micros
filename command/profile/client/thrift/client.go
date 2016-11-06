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

func (c client) AddProfile(jsonProfile string) string {
	reply, err := c.ProfileServiceClient.AddProfile(jsonProfile)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) UpdateProfile(profileID string, jsonProfile string) string {
	reply, err := c.ProfileServiceClient.UpdateProfile(profileID, jsonProfile)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) UpdateProfileStatus(profileID string, status bool) string {
	reply, err := c.ProfileServiceClient.UpdateProfileStatus(profileID, status)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) UpdateProfileBase(profileID string, mmap map[string]string) string {
	reply, err := c.ProfileServiceClient.UpdateProfileBase(profileID, mmap)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) UpdateProfileAgencyMembers(profileID string, agencyMembers string) string {
	reply, err := c.ProfileServiceClient.UpdateProfileAgencyMembers(profileID, agencyMembers)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}
