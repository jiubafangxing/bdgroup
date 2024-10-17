package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 加载路由设置
func LoadRoutes(r *gin.Engine) {
	// 定义一个 GET 路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// 定义一个 POST 路由，接收 JSON 数据
	r.POST("/setCookies", setCookies())
	r.POST("/swatchUser", swatchUser())
	r.GET("/groups", groups())
}
