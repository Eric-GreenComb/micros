package service

import ()

// CategoryService is the abstract representation of this service.
type CategoryService interface {
	Ping() string
	// Parameters:
	//  - Path
	LoadCategory(path string) bool

	// return: []category marshal json string
	GetCategories() string
	// Parameters:
	//  - Serialnumber
	// return: []subcategory marshal json string
	GetSubCategories(serialnumber int32) string
}
