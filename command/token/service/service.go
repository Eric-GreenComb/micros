package service

import ()

// TokenService is the abstract representation of this service.
type TokenService interface {
	Ping() string
	// Parameters:
	//  - Key
	//  - Ttype
	CreateToken(key string, ttype int64) string
	// Parameters:
	//  - Key
	//  - Ttype
	DeleteToken(key string, ttype int64) bool
}
