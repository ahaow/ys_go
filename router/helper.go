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
		HelperRouter.GET("/list", helper.List)      // 活动集列表
		HelperRouter.POST("/update", helper.Update) // 活动集更新
		HelperRouter.POST("/delete", helper.Delete) // 活动集删除
	}
}
