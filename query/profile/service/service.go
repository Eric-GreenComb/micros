package service

import ()

// Service is the abstract representation of this service.
type ProfileService interface {
	Ping() string
	// Parameters:
	//  - ProfileID
	GetProfile(profile_id string) string
	// Parameters:
	//  - UserID
	GetProfilesByUserId(user_id string) string
	// Parameters:
	//  - CategoryID
	//  - Timestamp
	//  - Pagesize
	GetProfilesByCategory(category_id int64, timestamp int64, pagesize int64) string
	// Parameters:
	//  - SubcategoryID
	//  - Timestamp
	//  - Pagesize
	GetProfilesBySubCategory(subcategory_id int64, timestamp int64, pagesize int64) string
	// Parameters:
	//  - OptionMmap
	//  - KeyMmap
	//  - Timestamp
	//  - Pagesize
	SearchProfiles(option_mmap map[string]int64, key_mmap map[string]string, timestamp int64, pagesize int64) string
}
