package thrift

import (
	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/command/workhistory/service"
	thriftworkhistory "github.com/banerwai/micros/command/workhistory/thrift/gen-go/workhistory"
)

// New returns an AddService that's backed by the Thrift client.
func New(cli *thriftworkhistory.WorkHistoryServiceClient, logger log.Logger) service.WorkHistoryService {
	return &client{cli, logger}
}

type client struct {
	*thriftworkhistory.WorkHistoryServiceClient
	log.Logger
}

func (c client) Ping() string {
	reply, err := c.WorkHistoryServiceClient.Ping()
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) UpdateWorkHistory(profileID, jsonWorkhistory string) string {
	reply, err := c.WorkHistoryServiceClient.UpdateWorkHistory(profileID, jsonWorkhistory)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}
