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

func (c client) GetAccountByUserID(userID string) string {
	reply, err := c.AccountServiceClient.GetAccountByUserID(userID)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) GetBillingByID(id string) string {
	reply, err := c.AccountServiceClient.GetBillingByID(id)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) GetDealBillingByUserID(userID string, timestamp int64, pagesize int64) string {
	reply, err := c.AccountServiceClient.GetDealBillingByUserID(userID, timestamp, pagesize)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) GetBillingByUserID(userID string, timestamp int64, pagesize int64) string {
	reply, err := c.AccountServiceClient.GetBillingByUserID(userID, timestamp, pagesize)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}
