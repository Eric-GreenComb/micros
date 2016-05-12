package thrift

import (
	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/query/resume/service"
	thriftresume "github.com/banerwai/micros/query/resume/thrift/gen-go/resume"
)

// New returns an AddService that's backed by the Thrift client.
func New(cli *thriftresume.ResumeServiceClient, logger log.Logger) service.ResumeService {
	return &client{cli, logger}
}

type client struct {
	*thriftresume.ResumeServiceClient
	log.Logger
}

func (c client) Ping() string {
	reply, err := c.ResumeServiceClient.Ping()
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) GetResume(id string) string {
	reply, err := c.ResumeServiceClient.GetResume(id)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}
