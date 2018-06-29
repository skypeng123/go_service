package routers

import (
	"go_service/services"
	"gopkg.in/gin-gonic/gin.v1"
)

func SetupRouter() *gin.Engine {
	//r := gin.Default()

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.NoRoute(services.Handle404)

	// Ping test
	r.GET("/ping", services.Ping)

	// 获取TOKEN
	r.POST("/token", services.Token)

	v1 := r.Group("/v1")
	v1.Use(services.AuthRequired())
	{
		//收费计算
		v1.POST("/charge", services.Charge)
	}

	return r
}
