/**
* @作者: 杨腾
* @date: 2022/6/23
* @describe: 用于创建用户的实体类
**/
package entity

import "gorm.io/gorm"

type (
	User struct {
		gorm.Model `json:"gorm_._model" form:"gorm_._model" query:"gorm_._model" uri:"gorm_._model"`
		// 这个是个性化的名字随便你取
		Name     string `gorm:"not null" json:"name" form:"name" query:"name" validate:"required"`
		Password string `gorm:"not null" json:"password" form:"password" query:"password" validate:"required"`
		// 主要通过邮箱判断用户是否存在
		Email string `gorm:"not null" json:"email" form:"email" query:"email"`
		// 用户头像
		HeadPortrait string `json:"head_portrait" form:"head_portrait" query:"head_portrait"`
		// 用户权限
		RoleID int `json:"role_id" form:"role_id" query:"role_id"`
		// 签名
		Motto string `json:"motto" form:"motto" query:"motto"`
		// 性别int类型默认为零
		Sex int `json:"sex" form:"sex" query:"sex"`
		// 关联的人物卡片
		CharacterID uint `gorm:"not null" json:"character_id" form:"character_id" query:"character_id" uri:"character_id"`
	}

	// Role 用户权限的实体类，权限系统采用三层架构
	// 角色表
	Role struct {
		gorm.Model  `json:"gorm_._model" form:"gorm_._model" query:"gorm_._model" uri:"gorm_._model"`
		Name        string `json:"name" form:"name" query:"name"`
		Description string `json:"description" form:"description" query:"description"`
	}

	// UserRole 用户角色表
	UserRole struct {
		gorm.Model `json:"gorm_._model" form:"gorm_._model" query:"gorm_._model" uri:"gorm_._model"`
		UserId     uint `json:"user_id" form:"user_id" query:"user_id"`
		RoleId     uint `json:"role_id" form:"role_id" query:"role_id"`
	}

	// Jur 权限表
	Jur struct {
		gorm.Model  `json:"gorm_._model" form:"gorm_._model" query:"gorm_._model" uri:"gorm_._model"`
		Name        string `json:"name" form:"name" query:"name"`
		Description string `json:"description" form:"description" query:"description"`
	}

	// JurRole 权限角色表
	JurRole struct {
		gorm.Model
		JurId  uint `json:"jur_id" form:"jur_id" query:"jur_id"`
		RoleId uint `json:"role_id" form:"role_id" query:"role_id"`
	}
)
