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

func (c client) RenderHello(tmpl, name string) string {
	reply, err := c.RenderServiceClient.RenderHello(tmpl, name)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}
