package bean

import ()

type Category struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	Subcategories []SubCategory `json:"subcategories"`
}

// Attributes:
//  - ID
//  - Name
type SubCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
