package category

import "gorm.io/gorm"

type (
	Category struct {
		gorm.Model `json:"gorm_._model" form:"gorm_._model" query:"gorm_._model" uri:"gorm_._model"`
		Name       string `json:"name,omitempty" form:"name" query:"name" uri:"name"`
		Children   uint   `json:"children,omitempty" form:"children" query:"children" uri:"children"`
		Describe   string `json:"describe,omitempty" form:"describe" query:"describe" uri:"describe"`
		Type       int    `json:"type" form:"type" query:"type" uri:"type"`
	}
)
