package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// 连接到 MySQL 数据库
	db, err := sql.Open("mysql", "user:password@tcp(db:3306)/mydb")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	defer db.Close()

	//// 创建表
	//_, err = db.Exec(`
	//	CREATE TABLE IF NOT EXISTS users (
	//		id INT AUTO_INCREMENT PRIMARY KEY,
	//		name VARCHAR(255),
	//		email VARCHAR(255)
	//	)
	//`)
	//if err != nil {
	//	fmt.Fprintf(err)
	//}

	// 插入数据
	_, err = db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "John Doe", "john@example.com")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	// 查询数据
	rows, err := db.Query("SELECT name, email FROM users")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var name, email string
		if err := rows.Scan(&name, &email); err != nil {
			fmt.Fprintf(w, err.Error())
		}
		fmt.Printf("Name: %s, Email: %s\n", name, email)
	}

	// ... 其他操作 ...
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
