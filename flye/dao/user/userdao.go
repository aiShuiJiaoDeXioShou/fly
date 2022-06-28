/**
* @作者:
* @date: 2022/6/23
* @describe: 用户管理数据库中用户表的增删改查
**/
package userdao

import (
	"flye/dao"
	"flye/entity"
	"log"
)

// 更新用户实体类在数据库里面的数据
func init() {
	err := dao.DB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Println("更新用户实体失败！")
	}
}

// 创建一个用户对象
