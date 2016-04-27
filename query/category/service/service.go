package service

import ()

// CategoryService is the abstract representation of this service.
type CategoryService interface {
	// Parameters:
	//  - Name
	SayHi(name string) string
	// Parameters:
	//  - ID
	// return: subcategory marshal json string
	GetDemoSubCategory(id string) string

	// Parameters:
	//  - CategoryID
	// return: []subcategory marshal json string
	GetDemoSubCategories(category_id string) string
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
