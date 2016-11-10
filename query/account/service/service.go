package service

import ()

// Service is the abstract representation of this service.
type AccountService interface {
	Ping() string
	// Parameters:
	//  - UserID
	GetAccountByUserID(userID string) string
	// Parameters:
	//  - ID
	GetBillingByID(ID string) string
	// Parameters:
	//  - UserID
	//  - Timestamp
	//  - Pagesize
	GetDealBillingByUserID(userID string, timestamp int64, pagesize int64) string
	// Parameters:
	//  - UserID
	//  - Timestamp
	//  - Pagesize
	GetBillingByUserID(userID string, timestamp int64, pagesize int64) string
}
