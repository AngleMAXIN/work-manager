package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// LoginAuthentication 登录验证
func LoginAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL
		fmt.Println(url.Path)
		if _, err := c.Cookie("wm_login"); err == nil {
			c.Next()
		} else {
			c.Abort()
			c.JSON(200, gin.H{
				"code": 301,
				"msg":  "用户未登录",
				"data": nil,
			})
		}

	}
}
