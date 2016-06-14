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

func (tb thriftBinding) CreateContact(json_contact string) (string, error) {
	r := tb.ContactService.CreateContact(json_contact)
	return r, nil
}

func (tb thriftBinding) ClientSignContact(contact_id string, status bool) (string, error) {
	r := tb.ContactService.ClientSignContact(contact_id, status)
	return r, nil
}

func (tb thriftBinding) FreelancerSignContact(contact_id string, status bool) (string, error) {
	r := tb.ContactService.FreelancerSignContact(contact_id, status)
	return r, nil
}

func (tb thriftBinding) DealContact(contact_id string, status bool) (string, error) {
	r := tb.ContactService.DealContact(contact_id, status)
	return r, nil
}

func (tb thriftBinding) UpdateContact(contact_id string, mmap map[string]string) (string, error) {
	r := tb.ContactService.UpdateContact(contact_id, mmap)
	return r, nil
}
