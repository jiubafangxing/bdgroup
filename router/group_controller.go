package router

import (
	"fmt"
	"github.com/bdgroup/cli/application"
	"github.com/bdgroup/service"
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
		} else {
			c.JSON(http.StatusOK, gin.H{
				"groups": shareGroups,
			})
		}

	}
}

func GetShareGroupFileList() func(c *gin.Context) {
	return func(c *gin.Context) {
		gid := c.Query("gid")
		files, err := application.FileLibraries(gid)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"files": files,
			})
		}
	}
}
func GetShareFiles() func(c *gin.Context) {
	return func(c *gin.Context) {
		var json service.ShareInfoParam
		if err := c.ShouldBindJSON(&json); err == nil {
			files, err := application.GetShareFiles(json)
			if err != nil {
				c.JSON(http.StatusBadGateway, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"files": files,
				})
			}
		} else {
			fmt.Println(err)
		}
	}
}
