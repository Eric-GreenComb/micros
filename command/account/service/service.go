package service

import ()

// AccountService Service is the abstract representation of this service.
type AccountService interface {
	Ping() string
	// Parameters:
	//  - JSONAccount
	CreateAccount(jsonAccount string) string
	// Parameters:
	//  - JSONBilling
	CreateBilling(jsonBilling string) string
	// Parameters:
	//  - BillingID
	DealBilling(billingID string) string
	// Parameters:
	//  - BillingID
	CancelBilling(billingID string) string
	// Parameters:
	//  - UserID
	GenAccount(userID string) string
}
