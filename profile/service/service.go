package service

import (
	thriftprofile "github.com/banerwai/micros/profile/thrift/gen-go/profile"
)

// Service is the abstract representation of this service.
type ProfileService interface {
	// Parameters:
	//  - ProfileID
	GetProfile(profile_id string) string

	// Parameters:
	//  - ProfileSearchCondition
	SearchProfiles(profile_search_condition *thriftprofile.ProfileSearchCondition, timestamp int64, pagesize int64) string
}
