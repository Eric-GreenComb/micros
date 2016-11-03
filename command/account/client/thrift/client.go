package thrift

import (
	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/command/account/service"
	thriftaccount "github.com/banerwai/micros/command/account/thrift/gen-go/account"
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

func (c client) CreateAccount(jsonAccount string) string {
	reply, err := c.AccountServiceClient.CreateAccount(jsonAccount)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) CreateBilling(jsonBilling string) string {
	reply, err := c.AccountServiceClient.CreateBilling(jsonBilling)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) DealBilling(billingID string) string {
	reply, err := c.AccountServiceClient.DealBilling(billingID)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) CancelBilling(billingID string) string {
	reply, err := c.AccountServiceClient.CancelBilling(billingID)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) GenAccount(userID string) string {
	reply, err := c.AccountServiceClient.GenAccount(userID)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}
