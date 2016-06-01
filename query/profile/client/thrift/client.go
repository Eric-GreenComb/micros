package thrift

import (
	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/query/profile/service"
	thriftprofile "github.com/banerwai/micros/query/profile/thrift/gen-go/profile"
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

func (c client) GetProfile(profile_id string) string {
	reply, err := c.ProfileServiceClient.GetProfile(profile_id)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) GetProfilesByUserId(user_id string) string {
	reply, err := c.ProfileServiceClient.GetProfilesByUserId(user_id)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) GetProfilesByCategory(category_id int64, timestamp int64, pagesize int64) string {
	reply, err := c.ProfileServiceClient.GetProfilesByCategory(category_id, timestamp, pagesize)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) GetProfilesBySubCategory(subcategory_id int64, timestamp int64, pagesize int64) string {
	reply, err := c.ProfileServiceClient.GetProfilesBySubCategory(subcategory_id, timestamp, pagesize)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) SearchProfiles(option_mmap map[string]int64, key_mmap map[string]string, timestamp int64, pagesize int64) string {
	reply, err := c.ProfileServiceClient.SearchProfiles(option_mmap, key_mmap, timestamp, pagesize)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}
