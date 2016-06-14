package main

import (
	"github.com/banerwai/micros/query/contact/service"
)

type thriftBinding struct {
	service.ContactService
}

func (tb thriftBinding) Ping() (string, error) {
	return tb.ContactService.Ping(), nil
}

func (tb thriftBinding) GetContactTpl(tplname string) (string, error) {
	return tb.ContactService.GetContactTpl(tplname), nil
}

func (tb thriftBinding) GetContact(contactid string) (string, error) {
	return tb.ContactService.GetContact(contactid), nil
}

func (tb thriftBinding) GetContactSignStatus(contactid string) (string, error) {
	return tb.ContactService.GetContactSignStatus(contactid), nil
}

func (tb thriftBinding) GetClientContact(clientemail string) (string, error) {
	return tb.ContactService.GetClientContact(clientemail), nil
}

func (tb thriftBinding) GetFreelancerContact(freelanceremail string) (string, error) {
	return tb.ContactService.GetFreelancerContact(freelanceremail), nil
}
