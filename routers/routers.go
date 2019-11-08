package routers

import (
	_ "work-manager/docs"
	v1 "work-manager/routers/api/v1"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	DebugMode := "debug"
	gin.SetMode(DebugMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/register", v1.Register)

		apiv1.POST("/login", v1.Login)

	}
	return r
}
