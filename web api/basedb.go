package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

// 连接数据库
func InitDB() {
	db, err := sql.Open("mysql", "root14:nH5QmLpwRKK8Jc6q@tcp(mysql2.sqlpub.com:3307)/root01")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(25)                 // 设置最大打开的连接数
	db.SetMaxIdleConns(10)                 // 设置最大空闲连接数
	db.SetConnMaxLifetime(5 * time.Minute) // 设置连接的最大生命周期
	defer db.Close()
	// 创建 users 表，添加 role 字段
	query := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE,
		password TEXT,
		role TEXT DEFAULT 'user'
	);`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database initialized.")
}

// 获取用户信息
func GetUser(db *sql.DB, username string) (string, string, error) {
	var password, role string
	err := db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&password)
	db.QueryRow("SELECT role FROM users WHERE username = ?", username).Scan(&role)
	return password, role, err
}
