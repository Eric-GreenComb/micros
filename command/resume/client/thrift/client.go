package thrift

import (
	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/command/resume/service"
	thriftresume "github.com/banerwai/micros/command/resume/thrift/gen-go/resume"
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

func (c client) AddResume(json_resume string) string {
	reply, err := c.ResumeServiceClient.AddResume(json_resume)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) UpdateResume(json_resume string) string {
	reply, err := c.ResumeServiceClient.UpdateResume(json_resume)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}
