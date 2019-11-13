package routers

import (
	"encoding/gob"
	_ "work-manager/docs"
	"work-manager/pkg/common"
	v1 "work-manager/routers/api/v1"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()
	r.MaxMultipartMemory = 8 << 20 // 8M

	// r.Use(middleware.LoginAuthentication())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gob.Register(common.UserBody{})
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("wm_login", store))

	DebugMode := "debug"
	gin.SetMode(DebugMode)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/register", v1.Register)
		apiv1.POST("/login", v1.Login)
		// apiv1.GET("/grade-info", handlers)
		apiv1.GET("/homework", v1.GetWorkList)
		apiv1.POST("/post-work", v1.PostHomeWork)
		apiv1.GET("/garde-homework", v1.GetGradeWorkList)
		apiv1.POST("/comment-work", v1.PutComment)

	}
	return r
}
