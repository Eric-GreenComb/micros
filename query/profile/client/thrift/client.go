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

func (c client) GetProfile(profileID string) string {
	reply, err := c.ProfileServiceClient.GetProfile(profileID)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) GetProfilesByUserID(userID string) string {
	reply, err := c.ProfileServiceClient.GetProfilesByUserID(userID)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) GetProfilesByCategory(categoryID int64, timestamp int64, pagesize int64) string {
	reply, err := c.ProfileServiceClient.GetProfilesByCategory(categoryID, timestamp, pagesize)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) GetProfilesBySubCategory(subcategoryID int64, timestamp int64, pagesize int64) string {
	reply, err := c.ProfileServiceClient.GetProfilesBySubCategory(subcategoryID, timestamp, pagesize)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) SearchProfiles(optionMap map[string]int64, keyMap map[string]string, timestamp int64, pagesize int64) string {
	reply, err := c.ProfileServiceClient.SearchProfiles(optionMap, keyMap, timestamp, pagesize)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}
