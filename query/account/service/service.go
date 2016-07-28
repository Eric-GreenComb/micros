package service

import ()

// Service is the abstract representation of this service.
type AccountService interface {
	Ping() string
	// Parameters:
	//  - UserID
	GetAccountByUserId(user_id string) string
	// Parameters:
	//  - ID
	GetBillingById(id string) string
	// Parameters:
	//  - UserID
	//  - Timestamp
	//  - Pagesize
	GetDealBillingByUserId(user_id string, timestamp int64, pagesize int64) string
	// Parameters:
	//  - UserID
	//  - Timestamp
	//  - Pagesize
	GetBillingByUserId(user_id string, timestamp int64, pagesize int64) string
}
