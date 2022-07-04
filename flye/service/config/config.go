/**
* @作者: 林河
* @date: 2022/7/4
* @describe:
**/
package config

import "github.com/gin-gonic/gin"

// 读取指定位置的配置文件，然后发送给前端，前端每当修改该配置文件的json数据，都会通过该接口发送给后端
// 实现时时修改这个配置文件
func NewConfigService(rt *gin.RouterGroup) {
	{
		config := rt.Group("/config")
		// 读取初始化设置
		config.GET("/int", intConfig)
	}
}

func intConfig(ctx *gin.Context) {

}