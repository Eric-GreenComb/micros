package service

import ()

// Service is the abstract representation of this service.
type ContactService interface {
	Ping() string
	// Parameters:
	//  - JSONContact
	CreateContact(json_contact string) string
	// Parameters:
	//  - ContactID
	//  - Status
	ClientSignContact(contact_id string, status bool) string
	// Parameters:
	//  - ContactID
	//  - Status
	FreelancerSignContact(contact_id string, status bool) string
	// Parameters:
	//  - ContactID
	//  - Status
	DealContact(contact_id string, status bool) string
	// Parameters:
	//  - ContactID
	//  - Mmap
	UpdateContact(contact_id string, mmap map[string]string) string
}
