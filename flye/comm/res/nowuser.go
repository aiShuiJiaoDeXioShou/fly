package res

import (
	"flye/entity"
	"github.com/gin-gonic/gin"
)

func NowUser(ctx *gin.Context) *entity.User {
	// 获取当前的用户对象
	value, exists := ctx.Get("current_user")
	if !exists {
		return nil
	}
	user := value.(*entity.User)
	return user
}
