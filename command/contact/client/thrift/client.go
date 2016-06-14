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

func (c client) CreateContact(json_contact string) string {
	reply, err := c.ContactServiceClient.CreateContact(json_contact)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) ClientSignContact(contact_id string, status bool) string {
	reply, err := c.ContactServiceClient.ClientSignContact(contact_id, status)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) FreelancerSignContact(contact_id string, status bool) string {
	reply, err := c.ContactServiceClient.FreelancerSignContact(contact_id, status)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) DealContact(contact_id string, status bool) string {
	reply, err := c.ContactServiceClient.DealContact(contact_id, status)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) UpdateContact(contact_id string, mmap map[string]string) string {
	reply, err := c.ContactServiceClient.UpdateContact(contact_id, mmap)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}
