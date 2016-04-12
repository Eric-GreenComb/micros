package service

import ()

// Service is the abstract representation of this service.
type ProfileService interface {
	// Parameters:
	//  - ID
	GetProfile(id string) string
	// Parameters:
	//  - JSONSearch
	//  - Timestamp
	//  - Pagesize
	SearchProfiles(json_search string, timestamp int64, pagesize int64) string
}
