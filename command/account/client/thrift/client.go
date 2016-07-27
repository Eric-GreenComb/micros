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

func (c client) CreateAccount(json_account string) string {
	reply, err := c.AccountServiceClient.CreateAccount(json_account)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) CreateBilling(json_billing string) string {
	reply, err := c.AccountServiceClient.CreateBilling(json_billing)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) DealBilling(billing_id string) string {
	reply, err := c.AccountServiceClient.DealBilling(billing_id)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) CancelBilling(billing_id string) string {
	reply, err := c.AccountServiceClient.CancelBilling(billing_id)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) GenAccount(user_id string) string {
	reply, err := c.AccountServiceClient.GenAccount(user_id)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}
