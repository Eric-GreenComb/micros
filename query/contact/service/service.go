package service

import ()

// ContactService is the abstract representation of this service.
type ContactService interface {
	Ping() string
	// Parameters:
	//  - Tplname
	GetContactTpl(tplname string) string
	// Parameters:
	//  - Contactid
	GetContact(contactid string) string
	// Parameters:
	//  - Contactid
	GetContactSignStatus(contactid string) string
	// Parameters:
	//  - Clientemail
	GetClientContact(clientemail string) string
	// Parameters:
	//  - Freelanceremail
	GetFreelancerContact(freelanceremail string) string
}
