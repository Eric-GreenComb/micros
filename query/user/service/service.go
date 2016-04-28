package service

import ()

// Service is the abstract representation of this service.
type UserService interface {
	Ping() string
	// Parameters:
	//  - Email
	GetUser(email string) string
	CountUser() int64
}
