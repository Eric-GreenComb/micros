package service

import ()

// TokenService is the abstract representation of this service.
type TokenService interface {
	Ping() string
	// Parameters:
	//  - Token
	//  - Ttype
	//  - Overhour
	VerifyToken(token string, ttype int64, overhour int64) int64
}
