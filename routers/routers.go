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
		apiv1.GET("/logout", v1.Logout)
		// 班级信息
		apiv1.GET("/grade-info", v1.GradeList)

		// 首页作业列表
		apiv1.GET("/homework", v1.GetWorkList)
		// 学生提交作业
		apiv1.POST("/post-work", v1.PostHomeWork)
		// 删除提交的作业
		// apiv1.DELETE("/delete-work", v1.DeleteWork)
		// 老师创建作业
		apiv1.POST("/create-homework", v1.CreateHomeWork)

		// 获取当前作业
		apiv1.GET("/workfile", v1.GetWorkFile)
		// 获取一个作业的所有提交
		apiv1.GET("/grade-homework", v1.GetGradeWorkList)
		// 老师批改作业
		apiv1.POST("/comment-work", v1.PutComment)

	}
	return r
}
