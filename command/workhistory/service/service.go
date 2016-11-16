package service

import ()

// WorkHistoryService Service is the abstract representation of this service.
type WorkHistoryService interface {
	Ping() string
	// Parameters:
	//  - ProfileID
	//  - JSONWorkhistory
	UpdateWorkHistory(profileID string, jsonWorkhistory string) string
}
