package service

import ()

// Service is the abstract representation of this service.
type ProfileService interface {
	Ping() string
	// Parameters:
	//  - ID
	GetProfile(id string) string
	// Parameters:
	//  - Email
	GetProfilesByEmail(email string) string
	// Parameters:
	//  - option_mmap
	//  - KeyMmap
	//  - Timestamp
	//  - Pagesize
	SearchProfiles(option_mmap map[string]int64, key_mmap map[string]string, timestamp int64, pagesize int64) string
}
