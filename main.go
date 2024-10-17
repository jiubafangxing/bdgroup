package main

import (
	"database/sql"
	"fmt"
	"github.com/bdgroup/config"
	"github.com/bdgroup/pkg/logger"
	"github.com/bdgroup/router"
	"github.com/gin-gonic/gin"
	"log/slog"
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

}

func main() {
	// 创建一个默认的路由引擎
	app := gin.Default()
	router.LoadRoutes(app)

	// 启动 HTTP 服务，监听在 8080 端口
	app.Run(":8080")
}
