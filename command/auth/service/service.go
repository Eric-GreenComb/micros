package service

import ()

// Service is the abstract representation of this service.
type AuthService interface {
	// Parameters:
	//  - Email
	//  - Pwd
	//  - FromUserId
	Register(email string, pwd string, fromUserId string) string
	// Parameters:
	//  - EmailOrUsername
	//  - Pwd
	Login(emailOrUsername string, pwd string) string
}
