package thrift

import (
	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/command/contact/service"
	thriftcontact "github.com/banerwai/micros/command/contact/thrift/gen-go/contact"
)

// New returns an AddService that's backed by the Thrift client.
func New(cli *thriftcontact.ContactServiceClient, logger log.Logger) service.ContactService {
	return &client{cli, logger}
}

type client struct {
	*thriftcontact.ContactServiceClient
	log.Logger
}

func (c client) Ping() string {
	reply, err := c.ContactServiceClient.Ping()
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) CreateContact(jsonContact string) string {
	reply, err := c.ContactServiceClient.CreateContact(jsonContact)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) ClientSignContact(contactID string, status bool) string {
	reply, err := c.ContactServiceClient.ClientSignContact(contactID, status)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) FreelancerSignContact(contactID string, status bool) string {
	reply, err := c.ContactServiceClient.FreelancerSignContact(contactID, status)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) DealContact(contactID string, status bool) string {
	reply, err := c.ContactServiceClient.DealContact(contactID, status)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) UpdateContact(contactID string, mmap map[string]string) string {
	reply, err := c.ContactServiceClient.UpdateContact(contactID, mmap)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}
