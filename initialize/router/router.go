package router

import (
	"time"
	"ys_go/global"
	"ys_go/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	port := global.Config.App.Port

	gin.SetMode("debug")

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"http://localhost:3000"}, // 你前端的地址，*表示全部允许
		AllowOrigins:     []string{"*"}, // 你前端的地址，*表示全部允许
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	g := r.Group("/api/v1")

	router.InitUserRouter(g)
	router.InitHelperRouter(g)

	if port == "" {
		port = ":3000"
	}

	r.Run(port)

}
