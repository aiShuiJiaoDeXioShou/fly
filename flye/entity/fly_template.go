/**
* @作者: 杨腾
* @date: 2022/6/23
* @describe: 这里面放的是人物模板的实体类
**/
package entity

import "gorm.io/gorm"

type (

	// FlyTemplate 通用模板的实体类
	FlyTemplate struct {
		// 首先创建时间id，创建时间这些东西
		gorm.Model `json:"gorm_._model" form:"gorm_._model" query:"gorm_._model" uri:"gorm_._model"`
		// 无论是物品还是人物，他都会有一个绘图，一个名字
		Name  string `json:"name,omitempty" form:"name" query:"name" uri:"name"`
		Image string `json:"image,omitempty" form:"image" query:"image" uri:"image"`
		// 还有一段基础描述
		Describe string `json:"describe,omitempty" form:"describe" query:"describe" uri:"describe"`
		// 版本控制
		Version string `json:"version,omitempty" form:"version" query:"version" uri:"version"`
		// 类型，这个是模板的类型，是人物还是物品或者是Base，玩家啥的，前端根据这个值来显示数据
		Type uint `json:"type,omitempty" form:"type" query:"type" uri:"type"`
	}

	// FlyTemplateField 通用模板单元格，这里记录扩展之后的单元格（基础之上的扩展单元格）
	// 写查询的时候可以通过关联父ID拿到子类的所有数据，然后再通过Data表拿到所有的数据，实现对于一张Cart数据的整体查询
	FlyTemplateField struct {
		gorm.Model `json:"gorm_._model" form:"gorm_._model" query:"gorm_._model" uri:"gorm_._model"`
		// 单元格的名字
		Name string `json:"name,omitempty" form:"name" query:"name" uri:"name"`
		// 单元格的类型
		Type string `json:"type,omitempty" form:"type" query:"type" uri:"type"`
		// 单元格的表达式，前端解析表达式,通过DSL显示数据
		Expression string `json:"expression,omitempty" form:"expression" query:"expression" uri:"expression"`
		// 单元格的默认值,父类单元格ID没有值,这里面的Value是他的默认值,这个是模板表,模板是用来创建其他模板的，注意这个Value是一个json数据
		DefValue string `json:"def_value,omitempty" form:"def_value" query:"def_value" uri:"def_value"`
		// Value的值在有数据的时候，默认显示Value的值
		Value string `json:"value,omitempty" form:"value" query:"value" uri:"value"`
		// 数据的类型，可选为一个数组，一个对象，一个字符串，一组数字，甚至一张图片，总之是一个json数据
		ValueType string `json:"value_type,omitempty" form:"value_type" query:"value_type" uri:"value_type"`
		// 所关联的父类单元格的ID,父类ID为0的话，就是根类
		FatherFieldID uint `json:"father_field_id" form:"father_field_id" query:"father_field_id" uri:"father_field_id"`
	}

	// FlyTemplateData 单元格数据表,用于存放通用模板表的数据
	FlyTemplateData struct {
		gorm.Model `json:"gorm_._model" form:"gorm_._model" query:"gorm_._model" uri:"gorm_._model"`
		// 绑定模板表的值
		FlyTemplateID uint `json:"fly_template_id" form:"fly_template_id" query:"fly_template_id" uri:"fly_template_id"`
		// 绑定单元格表的值
		FlyTemplateFieldID uint `json:"fly_template_field_id" form:"fly_template_field_id" query:"fly_template_field_id" uri:"fly_template_field_id"`
	}
)
