package service

import ()

// Service is the abstract representation of this service.
type ProfileService interface {
	// Parameters:
	//  - ProfileID
	GetProfile(profile_id string) string
	// Parameters:
	//  - Name
	GetProfileByCat(name string) string
	// Parameters:
	//  - Name
	GetProfileBySubCat(name string) string
}
