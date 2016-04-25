package thrift

import (
	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/command/token/service"
	thrifttoken "github.com/banerwai/micros/command/token/thrift/gen-go/token"
)

// New returns an AddService that's backed by the Thrift client.
func New(cli *thrifttoken.TokenServiceClient, logger log.Logger) service.TokenService {
	return &client{cli, logger}
}

type client struct {
	*thrifttoken.TokenServiceClient
	log.Logger
}

func (c client) NewToken_(key string, ttype int64) string {
	reply, err := c.TokenServiceClient.NewToken_(key, ttype)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) DeleteToken(key string, ttype int64) bool {
	reply, err := c.TokenServiceClient.DeleteToken(key, ttype)
	if err != nil {
		c.Logger.Log("err", err)
		return false
	}
	return reply
}
