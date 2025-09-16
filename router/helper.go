package router

import (
	"ys_go/api/helper"

	"github.com/gin-gonic/gin"
)

func InitHelperRouter(Router *gin.RouterGroup) {
	HelperRouter := Router.Group("helper")
	{
		HelperRouter.POST("/create", helper.Create) // 活动集创建
		HelperRouter.GET("/info/:id", helper.Info)  // 活动集详情
	}
}
