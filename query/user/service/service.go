package service

import ()

// Service is the abstract representation of this service.
type UserService interface {
	Ping() string
	// Parameters:
	//  - Email
	GetUserByEmail(email string) string
	// Parameters:
	//  - ID
	GetUserByID(id string) string
	CountUser() int64
}
