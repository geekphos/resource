package v1

import "gorm.io/datatypes"

type AllCategoryTagResponse struct {
	Name     string   `json:"name"`
	Children []string `json:"children"`
}

type CategoryTag struct {
	Category string
	Tags     datatypes.JSON
}
