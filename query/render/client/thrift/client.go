package thrift

import (
	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/query/render/service"
	thriftrender "github.com/banerwai/micros/query/render/thrift/gen-go/render"
)

// New returns an AddService that's backed by the Thrift client.
func New(cli *thriftrender.RenderServiceClient, logger log.Logger) service.RenderService {
	return &client{cli, logger}
}

type client struct {
	*thriftrender.RenderServiceClient
	log.Logger
}

func (c client) Ping() string {
	reply, err := c.RenderServiceClient.Ping()
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) RenderTpl(tplname string, key_mmap map[string]string) string {
	reply, err := c.RenderServiceClient.RenderTpl(tplname, key_mmap)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}
