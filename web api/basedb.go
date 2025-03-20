package main

import (
	"database/sql"
	"fmt"
	"log"
)

// 连接数据库
func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}

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
	return db
}

// 获取用户信息
func GetUser(db *sql.DB, username string) (string, string, error) {
	var password, role string
	err := db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&password)
	db.QueryRow("SELECT role FROM users WHERE username = ?", username).Scan(&role)
	return password, role, err
}
