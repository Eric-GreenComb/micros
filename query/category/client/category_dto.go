package main

import ()

// Attributes:
//  - ID
//  - Name
//  - Desc
//  - Subcategories
type Category struct {
	ID            string        `json:"id"`
	SerialNumber  int32         `json:"serialnumber"`
	Name          string        `json:"name"`
	Desc          string        `json:"desc"`
	Fa            string        `json:"fa"`
	Subcategories []SubCategory `json:"subcategories"`
}

// Attributes:
//  - ID
//  - Name
//  - Desc
type SubCategory struct {
	ID           string `json:"id"`
	SerialNumber int32  `json:"serialnumber"`
	Name         string `json:"name"`
	Desc         string `json:"desc"`
}
