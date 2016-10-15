package service

import ()

// WorkHistoryService is the abstract representation of this service.
type WorkHistoryService interface {
	Ping() string
	// Parameters:
	//  - ProfileID
	GetWorkHistory(profileID string) string
}
