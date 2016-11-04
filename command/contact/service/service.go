package service

import ()

// ContactService Service is the abstract representation of this service.
type ContactService interface {
	Ping() string
	// Parameters:
	//  - JSONContact
	CreateContact(jsonContact string) string
	// Parameters:
	//  - ContactID
	//  - Status
	ClientSignContact(contactID string, status bool) string
	// Parameters:
	//  - ContactID
	//  - Status
	FreelancerSignContact(contactID string, status bool) string
	// Parameters:
	//  - ContactID
	//  - Status
	DealContact(contactID string, status bool) string
	// Parameters:
	//  - ContactID
	//  - Mmap
	UpdateContact(contactID string, mmap map[string]string) string
}
