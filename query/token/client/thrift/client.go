package thrift

import (
	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/query/token/service"
	thrifttoken "github.com/banerwai/micros/query/token/thrift/gen-go/token"
)

// New returns an AddService that's backed by the Thrift client.
func New(cli *thrifttoken.TokenServiceClient, logger log.Logger) service.TokenService {
	return &client{cli, logger}
}

type client struct {
	*thrifttoken.TokenServiceClient
	log.Logger
}

func (c client) Ping() string {
	reply, err := c.TokenServiceClient.Ping()
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) VerifyToken(key string, ttype int64, overhour float64) int64 {
	reply, err := c.TokenServiceClient.VerifyToken(key, ttype, overhour)
	if err != nil {
		c.Logger.Log("err", err)
		return -1
	}
	return reply
}
