package service

import ()

// ProfileService Service is the abstract representation of this service.
type ProfileService interface {
	Ping() string
	// Parameters:
	//  - JSONProfile
	AddProfile(jsonProfile string) string
	// Parameters:
	//  - ProfileID
	//  - JSONProfile
	UpdateProfile(profileID string, jsonProfile string) string
	// Parameters:
	//  - ProfileID
	//  - Status
	UpdateProfileStatus(profileID string, status bool) string
	// Parameters:
	//  - ProfileID
	//  - Mmap
	UpdateProfileBase(profileID string, mmap map[string]string) string
	// Parameters:
	//  - ProfileID
	//  - AgencyMembers
	UpdateProfileAgencyMembers(profileID string, agencyMembers string) string
}
