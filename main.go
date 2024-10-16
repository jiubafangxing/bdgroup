package main

import (
	"database/sql"
	"fmt"
	"github.com/bdgroup/config"
	"github.com/bdgroup/pkg/logger"
	"github.com/bdgroup/service"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func init() {
	l := logger.SetLogLevel(os.Getenv("LOG_LEVEL"))
	slog.SetDefault(l)
	err := config.Instance.Init()
	if err != nil {
		fmt.Println(err)
	}
	//初始化数据库
	// 1. 打开或创建 SQLite 数据库
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	// 定义一个 GET 路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 定义一个 POST 路由，接收 JSON 数据
	r.POST("/setCookies", func(c *gin.Context) {
		var json struct {
			BDCookies string `json:"bdCookies" binding:"required"`
		}
		// 绑定 JSON 数据并校验
		if err := c.ShouldBindJSON(&json); err == nil {
			bdnInfo, err := config.Instance.SetBdnInfoByCookies(json.BDCookies)
			if err != nil {
				c.JSON(http.StatusBadGateway, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"username": bdnInfo.Username,
				})
				curtime := time.Now()
				if nil == bdnInfo {
					service.InsertUser(bdnInfo.Username, bdnInfo.UK, curtime, curtime, bdnInfo.Cookies)
				} else {
					service.UpdateUser(bdnInfo.UK, bdnInfo.Cookies, curtime)
				}
				log.Printf("百度网盘登录验证登录成功, 昵称:%s", bdnInfo.Username)
			}
			return
		} else {
			log.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{})
	})
	r.POST("/swatchUser", func(c *gin.Context) {
		var json struct {
			UserName string `json:"userName" binding:"required"`
		}
		if err := c.ShouldBindJSON(&json); err == nil {
			user, err := service.GetByUserName(json.UserName)
			if err == nil {
				bdnInfo, err := config.Instance.SetBdnInfoByCookies(user.Cookies)
				if err != nil {
					c.JSON(http.StatusBadGateway, gin.H{
						"error": err.Error(),
					})
				} else {
					c.JSON(http.StatusOK, gin.H{
						"username": bdnInfo.Username,
					})
					log.Printf("百度网盘登录验证登录成功, 昵称:%s", bdnInfo.Username)
				}
			}
		} else {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
		}
	})
	// 启动 HTTP 服务，监听在 8080 端口
	r.Run(":8080")
}
