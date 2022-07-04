/**
* @作者: 林河
* @date: 2022/7/4
* @describe:
**/
package config

import (
	"flye/comm/res"
	"github.com/ghodss/yaml"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

// 读取指定位置的配置文件，然后发送给前端，前端每当修改该配置文件的json数据，都会通过该接口发送给后端
// 实现时时修改这个配置文件
func NewConfigService(rt *gin.RouterGroup) {
	{
		config := rt.Group("/config")
		// 读取初始化设置
		config.GET("/init", intConfig)
	}
}

// 读取配置文件并且发送
func intConfig(ctx *gin.Context) {
	// 读取根源目录下的config.yaml 文件
	file, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(200, "路径错误：读取config.yaml文件失败，请确定文件地址是否正确\n"+err.Error())
		return
	}
	toJSON, err := yaml.YAMLToJSON(file)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(200, res.Error("josn解析错误：类型解析失败\n"+err.Error()))
		return
	}
	ctx.JSON(200, res.Ok(string(toJSON)))
	// 前端每当修改该配置文件的json数据，都会通过该接口发送给后端
	// 实现时时修改这个配置文件
}
