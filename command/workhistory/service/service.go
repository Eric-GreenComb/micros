package service

import ()

// Service is the abstract representation of this service.
type WorkHistoryService interface {
	Ping() string
	// Parameters:
	//  - JSONWorkhistory
	AddWorkHistory(json_workhistory string) string
	// Parameters:
	//  - JSONWorkhistory
	UpdateWorkHistory(json_workhistory string) string
}
