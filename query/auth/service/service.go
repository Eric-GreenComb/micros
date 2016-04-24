package service

import ()

// Service is the abstract representation of this service.
type AuthService interface {
	// Parameters:
	//  - EmailOrUsername
	//  - Pwd
	Login(emailOrUsername string, pwd string) string
}
