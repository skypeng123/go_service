package routers

import (
	"go_service/services"
	"gopkg.in/gin-gonic/gin.v1"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		// Ping test
		v1.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong")
		})

		// 获取TOKEN
		v1.POST("/token", services.GetToken)
	}

	// 收费计算
	//r.POST("/charge", charge)

	return r
}
