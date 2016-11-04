package main

import (
	"github.com/banerwai/micros/command/contact/service"
)

type thriftBinding struct {
	service.ContactService
}

func (tb thriftBinding) Ping() (string, error) {
	r := tb.ContactService.Ping()
	return r, nil
}

func (tb thriftBinding) CreateContact(jsonContact string) (string, error) {
	r := tb.ContactService.CreateContact(jsonContact)
	return r, nil
}

func (tb thriftBinding) ClientSignContact(contactID string, status bool) (string, error) {
	r := tb.ContactService.ClientSignContact(contactID, status)
	return r, nil
}

func (tb thriftBinding) FreelancerSignContact(contactID string, status bool) (string, error) {
	r := tb.ContactService.FreelancerSignContact(contactID, status)
	return r, nil
}

func (tb thriftBinding) DealContact(contactID string, status bool) (string, error) {
	r := tb.ContactService.DealContact(contactID, status)
	return r, nil
}

func (tb thriftBinding) UpdateContact(contactID string, mmap map[string]string) (string, error) {
	r := tb.ContactService.UpdateContact(contactID, mmap)
	return r, nil
}
