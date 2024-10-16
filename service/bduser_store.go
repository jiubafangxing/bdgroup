package service

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// User 结构体代表数据库中的用户记录
type BDUser struct {
	ID            int
	Username      string
	UK            string
	CreateTime    string
	LastLoginTime string
	Cookies       string
}

// 全局变量db，用于数据库操作
var db *sql.DB

func init() {
	// 初始化数据库
	// 1. 打开或创建 SQLite 数据库
	var err error
	db, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	// 注意：不要在这里使用defer db.Close()，因为init函数只在程序启动时执行一次
	// defer db.Close()会在init函数结束时关闭数据库连接，这不是我们想要的

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
		log.Fatal(err)
	}
}

// GetUserByUK 查询指定uk的用户信息
func GetUserByUK(uk string) (*BDUser, error) {
	var user BDUser
	err := db.QueryRow(`SELECT id, username, uk, create_time, last_login_time, cookies FROM bd_user WHERE uk = ?;`, uk).Scan(&user.ID, &user.Username, &user.UK, &user.CreateTime, &user.LastLoginTime, &user.Cookies)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // 没有找到记录，返回nil
		} else {
			return nil, err // 发生错误，返回错误信息
		}
	}
	return &user, nil // 成功找到记录，返回User指针
}

// InsertUser 插入新用户
func InsertUser(username, uk, createTime, lastLoginTime, cookies string) (int64, error) {
	stmt, err := db.Prepare(`INSERT INTO bd_user (username, uk, create_time, last_login_time, cookies) VALUES (?, ?, ?, ?, ?);`)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(username, uk, createTime, lastLoginTime, cookies)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
