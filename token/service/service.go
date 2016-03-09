package service

import ()

// TokenService is the abstract representation of this service.
type TokenService interface {
	NewToken_(key string, ttype int64) string
	DeleteToken(key string, ttype int64) bool
	VerifyToken(key string, ttype int64) int64
}
