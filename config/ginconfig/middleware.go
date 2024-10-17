package ginconfig

import (
	"github.com/gin-gonic/gin"
)

// CORS中间件函数
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 添加Access-Control-Allow-Origin头部
		c.Header("Access-Control-Allow-Origin", "*")

		// 如果需要，还可以添加其他CORS相关的头部
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// 继续处理请求
		c.Next()
	}
}

func main() {
	// 创建一个Gin引擎
	r := gin.Default()

	// 使用CORS中间件
	r.Use(CORSMiddleware())

	// 设置路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 启动服务器
	r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
}
