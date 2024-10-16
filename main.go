package main

import (
	"database/sql"
	"fmt"
	"github.com/bdgroup/cli/cmds"
	"github.com/bdgroup/config"
	"github.com/bdgroup/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"log"
	"log/slog"
	"net/http"
	"os"
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

	// 2. 创建表
	createTableSQL := `CREATE TABLE IF NOT EXISTS bd_user (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		uk TEXT UNIQUE,
		create_time DATETIME DEFAULT CURRENT_TIMESTAMP,
		last_login_time DATETIME,
		cookies TEXT
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main2() {
	app := cmds.NewApp()
	app.Commands = []cli.Command{}
	app.Commands = append(app.Commands, cmds.NewLoginCommand()...)
	app.Commands = append(app.Commands, cmds.NewBdCommand()...)

	app.Action = cmds.DefaultAction

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
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
				log.Printf("百度网盘登录验证登录成功, 昵称:%s,bdstoken:%s", bdnInfo.Username, bdnInfo.Bdstoken)
			}
			return
		} else {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{})
	})

	// 启动 HTTP 服务，监听在 8080 端口
	r.Run(":8080")
}
