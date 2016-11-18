package service

import ()

// ProfileService is the abstract representation of this service.
type ProfileService interface {
	Ping() string
	// Parameters:
	//  - ProfileID
	GetProfile(profileID string) string
	// Parameters:
	//  - UserID
	GetProfilesByUserID(userID string) string
	// Parameters:
	//  - CategoryID
	//  - Timestamp
	//  - Pagesize
	GetProfilesByCategory(categoryID int64, timestamp int64, pagesize int64) string
	// Parameters:
	//  - SubcategoryID
	//  - Timestamp
	//  - Pagesize
	GetProfilesBySubCategory(subcategoryID int64, timestamp int64, pagesize int64) string
	// Parameters:
	//  - OptionMmap
	//  - KeyMmap
	//  - Timestamp
	//  - Pagesize
	SearchProfiles(optionMap map[string]int64, keyMap map[string]string, timestamp int64, pagesize int64) string
}
