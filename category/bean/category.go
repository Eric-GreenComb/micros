package bean

import (
// "encoding/json"
// "fmt"
)

type Category struct {
	ID            string        `json:"id"`
	Name          string        `json:"name,omitempty"`
	SubCategories []SubCategory `json:"subcategories,omitempty"`
}

type SubCategory struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

// func (self *CategoryBean) GetCategoryBeanJson() string {
// 	b, err := json.Marshal(&self)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	return string(b)
// }

// func (self *CategoryBean) GetCategories() []StringKeyValue {
// 	var _keyValues []StringKeyValue

// 	for _, _category := range self.Categories {
// 		var _keyValue StringKeyValue
// 		_keyValue.Key = _category.ID
// 		_keyValue.Value = _category.Name

// 		_keyValues = append(_keyValues, _keyValue)
// 	}

// 	return _keyValues
// }

// func (self *CategoryBean) GetCategoriesJson() string {
// 	_keyValues := self.GetCategories()
// 	if len(_keyValues) == 0 {
// 		fmt.Println("Categories is null")
// 		return ""
// 	}
// 	b, err := json.Marshal(_keyValues)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	return string(b)
// }

// func (self *CategoryBean) GetCategory(id string) (Category, bool) {
// 	var category Category

// 	bHas := false
// 	for _, _category := range self.Categories {
// 		if id != _category.ID {
// 			continue
// 		}

// 		category = _category
// 		bHas = true
// 		break
// 	}

// 	return category, bHas
// }

// func (self *CategoryBean) GetCategoryJson(id string) string {
// 	_category, _bHas := self.GetCategory(id)
// 	if !_bHas {
// 		fmt.Println("Category is null")
// 		return ""
// 	}
// 	b, err := json.Marshal(_category)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	return string(b)
// }

// func (self *CategoryBean) GetSubCategories(id string) ([]SubCategory, bool) {
// 	var _subcategories []SubCategory

// 	bHas := false
// 	for _, _category := range self.Categories {
// 		if id != _category.ID {
// 			continue
// 		}

// 		if len(_category.SubCategories) > 0 {
// 			bHas = true
// 			_subcategories = _category.SubCategories
// 		}

// 		break
// 	}

// 	return _subcategories, bHas
// }

// func (self *CategoryBean) GetSubCategoriesJson(id string) string {
// 	_subcategories, _bHas := self.GetSubCategories(id)
// 	if !_bHas {
// 		fmt.Println("SubCategories is null")
// 		return ""
// 	}
// 	b, err := json.Marshal(_subcategories)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	return string(b)
// }
