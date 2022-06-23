package main

import (
	"flye/dao"
	"flye/dao/category"
	"log"
	"testing"
)

// 测试数据库是否连接成功
func Test1(t *testing.T) {
	// 连接成功的话，添加一条数据
	err := dao.DB.Save(&category.Category{
		Name:     "",
		Children: 0,
		Describe: "第一次跑团人物卡",
		Type:     0,
	}).Error

	if err != nil {
		panic(err)
		return
	}

	// 查找指定字段
	var categorys []category.Category
	err = dao.DB.Find(&categorys).Error
	if err != nil {
		panic(err)
		return
	}
	// 如果找到该数据，就打印到控制台上面
	for _, c := range categorys {
		log.Println(c)
	}
	// 逻辑删除该条数据
	err = dao.DB.Where("id in (?)", []uint{1, 2, 3, 4}).Delete(&category.Category{}).Error
}
