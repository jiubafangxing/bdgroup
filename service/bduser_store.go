package service

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

// User 结构体代表数据库中的用户记录
type BDUser struct {
	ID            int
	Username      string
	UK            string
	CreateTime    time.Time
	LastLoginTime time.Time
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
		username TEXT NOT NULL,
		uk TEXT   PRIMARY KEY,
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
func InsertUser(username, uk string, createTime, lastLoginTime time.Time, cookies string) (int64, error) {
	stmt, err := db.Prepare(`INSERT INTO bd_user (username, uk, create_time, last_login_time, cookies) VALUES (?, ?, ?, ?, ?);`)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(username, uk, createTime.Format("2006-01-02 15:04:05"), lastLoginTime.Format("2006-01-02 15:04:05"), cookies)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// UpdateUser 更新用户信息
func UpdateUser(uk string, cookies string, lastLoginTime time.Time) error {
	// 准备 SQL 更新语句
	stmt, err := db.Prepare(`UPDATE bd_user SET cookies = ?, last_login_time = ? WHERE uk = ?;`)
	if err != nil {
		return err
	}
	// 执行更新操作
	_, err = stmt.Exec(cookies, lastLoginTime.Format("2006-01-02 15:04:05"), uk)
	if err != nil {
		return err
	}
	return nil
}

// 依据用户名查询用户
func GetByUserName(userName string) (*BDUser, error) {
	var user BDUser
	err := db.QueryRow(`SELECT id, username, uk, create_time, last_login_time, cookies FROM bd_user WHERE username = ?;`, userName).Scan(&user.ID, &user.Username, &user.UK, &user.CreateTime, &user.LastLoginTime, &user.Cookies)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // 没有找到记录，返回nil
		} else {
			return nil, err // 发生错误，返回错误信息
		}
	}
	return &user, nil // 成功找到记录，返回User指针
}
