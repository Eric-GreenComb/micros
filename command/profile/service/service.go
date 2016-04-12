package service

import ()

// Service is the abstract representation of this service.
type ProfileService interface {
	// Parameters:
	//  - JSONProfile
	AddProfile(json_profile string) string
	// Parameters:
	//  - JSONProfile
	UpdateProfile(json_profile string) string
	// Parameters:
	//  - ID
	DeleteProfile(id string) string
}
