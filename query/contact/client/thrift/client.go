package thrift

import (
	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/query/contact/service"
	thriftcontact "github.com/banerwai/micros/query/contact/thrift/gen-go/contact"
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

func (c client) GetContactTpl(tplname string) string {
	reply, err := c.ContactServiceClient.GetContactTpl(tplname)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) GetContact(contactid string) string {
	reply, err := c.ContactServiceClient.GetContact(contactid)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) GetContactSignStatus(contactid string) string {
	reply, err := c.ContactServiceClient.GetContactSignStatus(contactid)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) GetClientContact(clientemail string) string {
	reply, err := c.ContactServiceClient.GetClientContact(clientemail)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) GetFreelancerContact(freelanceremail string) string {
	reply, err := c.ContactServiceClient.GetFreelancerContact(freelanceremail)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}
