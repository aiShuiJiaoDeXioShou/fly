package entity

import "gorm.io/gorm"

type (
	Category struct {
		gorm.Model `json:"gorm_._model" form:"gorm_._model" query:"gorm_._model" uri:"gorm_._model"`
		Name       string `json:"name,omitempty" form:"name" query:"name" uri:"name"`
		Children   uint   `json:"children,omitempty" form:"children" query:"children" uri:"children"`
		Describe   string `json:"describe,omitempty" form:"describe" query:"describe" uri:"describe"`
		Image      string `json:"image,omitempty" form:"image" query:"image" uri:"image"`
		Type       int    `json:"type" form:"type" query:"type" uri:"type"`
	}

	// 通过关联表绑定Category表和FlyTemplate表
	CategoryToFlyTemplate struct {
		gorm.Model    `json:"gorm_._model" form:"gorm_._model" query:"gorm_._model" uri:"gorm_._model"`
		CategoryID    uint `json:"category_id,omitempty" form:"category_id" query:"category_id" uri:"category_id"`
		FlyTemplateID uint `json:"fly_template_id,omitempty" form:"fly_template_id" query:"fly_template_id" uri:"fly_template_id"`
	}
)
