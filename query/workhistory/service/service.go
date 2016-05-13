package service

import ()

// Service is the abstract representation of this service.
type WorkHistoryService interface {
	Ping() string
	// Parameters:
	//  - ProfileID
	GetWorkHistory(profile_id string) string
}
