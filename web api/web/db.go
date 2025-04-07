package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := db{
		dsn: "root14:nH5QmLpwRKK8Jc6q@tcp(mysql2.sqlpub.com:3307)/root01",
	}
	db.init()
	db.creatTable()

}

type db struct {
	dsn string
}

func (mysql db) init() (basedata sql.DB) {
	// 连接 MySQL 数据库
	db, err := sql.Open("mysql", mysql.dsn)
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	db.SetMaxOpenConns(25)                 // 设置最大打开的连接数
	db.SetMaxIdleConns(10)                 // 设置最大空闲连接数
	db.SetConnMaxLifetime(5 * time.Minute) // 设置连接的最大生命周期

	// 测试连接
	err = db.Ping()
	if err != nil {
		log.Fatal("数据库不可用:", err)
	}
	fmt.Println("成功连接 MySQL 数据库！")
	return
}
func (mysql db) creatTable() {
	db, _ := sql.Open("mysql", mysql.dsn)
	// 创建表的 SQL 语句
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		age INT NOT NULL,
		email VARCHAR(255) UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	// 执行 SQL 语句

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("创建表失败:", err)
	}

	fmt.Println("成功创建 users 表！")
}
func postDB(db *sql.DB, table string) {
	result, err := db.Exec("INSERT INTO ? (name, age) VALUES (?, ?)", table, "Alice", 30)
	if err != nil {
		log.Fatal("插入失败:", err)
	}
	id, _ := result.LastInsertId()
	fmt.Println("新插入的 ID:", id)

}
func updateDB(db *sql.DB) {
	_, err := db.Exec("UPDATE users SET age = ? WHERE name = ?", 35, "Alice")
	if err != nil {
		log.Fatal("更新失败:", err)
	}

}
func deleteDB(db *sql.DB) {
	_, err := db.Exec("DELETE FROM users WHERE name = ?", "Alice")
	if err != nil {
		log.Fatal("删除失败:", err)
	}

}
func getDB(db *sql.DB) {
	var name string
	err := db.QueryRow("SELECT name FROM users WHERE id = ?", 1).Scan(&name)
	if err != nil {
		log.Fatal("查询失败:", err)
	}
	fmt.Println("用户名称:", name)

	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal("查询失败:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

}
