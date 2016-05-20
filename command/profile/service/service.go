package service

import ()

// Service is the abstract representation of this service.
type ProfileService interface {
	Ping() string
	// Parameters:
	//  - JSONProfile
	AddProfile(json_profile string) string
	// Parameters:
	//  - ProfileID
	//  - JSONProfile
	UpdateProfile(profile_id string, json_profile string) string
	// Parameters:
	//  - ProfileID
	//  - Status
	UpdateProfileStatus(profile_id string, status bool) string
	// Parameters:
	//  - ProfileID
	//  - Mmap
	UpdateProfileBase(profile_id string, mmap map[string]string) string
	// Parameters:
	//  - ProfileID
	//  - AgencyMembers
	UpdateProfileAgencyMembers(profile_id string, agency_members string) string
}
