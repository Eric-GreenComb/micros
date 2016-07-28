package thrift

import (
	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/query/account/service"
	thriftaccount "github.com/banerwai/micros/query/account/thrift/gen-go/account"
)

// New returns an AddService that's backed by the Thrift client.
func New(cli *thriftaccount.AccountServiceClient, logger log.Logger) service.AccountService {
	return &client{cli, logger}
}

type client struct {
	*thriftaccount.AccountServiceClient
	log.Logger
}

func (c client) Ping() string {
	reply, err := c.AccountServiceClient.Ping()
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) GetAccountByUserId(user_id string) string {
	reply, err := c.AccountServiceClient.GetAccountByUserId(user_id)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) GetBillingById(id string) string {
	reply, err := c.AccountServiceClient.GetBillingById(id)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) GetDealBillingByUserId(user_id string, timestamp int64, pagesize int64) string {
	reply, err := c.AccountServiceClient.GetDealBillingByUserId(user_id, timestamp, pagesize)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) GetBillingByUserId(user_id string, timestamp int64, pagesize int64) string {
	reply, err := c.AccountServiceClient.GetBillingByUserId(user_id, timestamp, pagesize)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}
