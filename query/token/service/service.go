package service

import ()

// TokenService is the abstract representation of this service.
type TokenService interface {
	// Parameters:
	//  - Token
	//  - Ttype
	//  - Overhour
	VerifyToken(token string, ttype int64, overhour float64) int64
}
