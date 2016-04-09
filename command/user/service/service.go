package service

import ()

// Service is the abstract representation of this service.
type UserService interface {
	// Parameters:
	//  - Email
	//  - Usernameraw
	//  - Pwd
	CreateUser(email string, usernameraw string, pwd string) string
	// Parameters:
	//  - Email
	//  - Oldpwd
	//  - Newpwd_
	UpdatePwd(email string, oldpwd string, newpwd string) bool
	// Parameters:
	//  - Token
	ActiveUser(token string) bool
	CountUser() int64
}
