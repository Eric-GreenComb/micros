package service

import ()

// Service is the abstract representation of this service.
type AuthService interface {
	Ping() string
	// Parameters:
	//  - Email
	//  - Pwd
	Login(email string, pwd string) string
}
