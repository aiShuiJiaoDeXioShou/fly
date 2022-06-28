/**
* @作者: 林河
* @date: 2022/6/27
* @describe:
**/
package qidiansort

import (
	"flye/comm/reptile/qidian"
	"github.com/gin-gonic/gin"
)

var rq = qidian.NewQiDianRankingReptile(qidian.NOVEL_RANKING)

// 创建一个新的爬取起点中文网服务
func NewQiDianSort(rt *gin.RouterGroup) {
	{
		qd := rt.Group("/qidian")
		qd.GET("/yuep", yuep)
	}
}

// 获取月票榜的数据
func yuep(ctx *gin.Context) {
	yueps := rq.GetRanking()
	ctx.JSON(200, yueps)
}
