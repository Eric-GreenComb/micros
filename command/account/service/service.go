package service

import ()

// Service is the abstract representation of this service.
type AccountService interface {
	Ping() string
	// Parameters:
	//  - JSONAccount
	CreateAccount(json_account string) string
	// Parameters:
	//  - JSONBilling
	CreateBilling(json_billing string) string
	// Parameters:
	//  - BillingID
	DealBilling(billing_id string) string
	// Parameters:
	//  - BillingID
	CancelBilling(billing_id string) string
	// Parameters:
	//  - UserID
	GenAccount(user_id string) string
}
