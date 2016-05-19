package service

import ()

// Service is the abstract representation of this service.
type WorkHistoryService interface {
	Ping() string
	// Parameters:
	//  - ProfileID
	//  - JSONWorkhistory
	UpdateWorkHistory(profile_id string, json_workhistory string) string
}
