package service

import ()

// UserService is the abstract representation of this service.
type UserService interface {
	Ping() string
	// Parameters:
	//  - Email
	//  - Mmap
	CreateUser(mmap map[string]string) string
	// Parameters:
	//  - Email
	//  - Newpwd_
	ResetPwd(email string, newpwd string) bool
	// Parameters:
	//  - Email
	//  - Token
	ActiveUser(email string) bool
}
