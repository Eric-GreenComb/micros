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

func (c client) UpdateProfile(profile_id string, json_profile string) string {
	reply, err := c.ProfileServiceClient.UpdateProfile(profile_id, json_profile)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) UpdateProfileStatus(profile_id string, status bool) string {
	reply, err := c.ProfileServiceClient.UpdateProfileStatus(profile_id, status)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) UpdateProfileBase(profile_id string, mmap map[string]string) string {
	reply, err := c.ProfileServiceClient.UpdateProfileBase(profile_id, mmap)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) UpdateProfileAgencyMembers(profile_id string, agency_members string) string {
	reply, err := c.ProfileServiceClient.UpdateProfileAgencyMembers(profile_id, agency_members)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}
