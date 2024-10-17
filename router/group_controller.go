package router

import (
	"github.com/bdgroup/cli/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 查询指定用户名的群组
func groups() func(c *gin.Context) {
	return func(c *gin.Context) {
		shareGroups, err := application.ShareGroups()
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"groups": shareGroups,
		})

	}
}
