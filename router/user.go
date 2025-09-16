package router

import (
	"ys_go/api/user"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("/register", user.Register) // 用户注册
		UserRouter.POST("/login", user.Login)       // 用户登录
		UserRouter.POST("/update", user.Update)     // 用户更新
		UserRouter.GET("/list", user.List)          // 列表
	}
}
