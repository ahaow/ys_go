package middleware

import (
	"ys_go/global"
	"ys_go/response"
	"ys_go/utils/jwt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claims, err := jwt.ParseJWT(token)

	global.Log.Error("AuthMiddleware", err)
	if err != nil {
		// token失效
		c.JSON(0, response.Response{
			Code: 401,
			Data: nil,
			Msg:  "token认证失败",
		})
		c.Abort()
		return
	}
	// 将用户信息存到上下文，方便后续使用
	c.Set("user_id", claims.UserId)
	c.Set("user_name", claims.Username)
	c.Next()
}
