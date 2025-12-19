package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var SQLClient *sql.DB

func InitMySQL() {
	var err error
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "root:@tcp(127.0.0.1:3306)/guvi_db?parseTime=true"
	}
	SQLClient, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to open MySQL:", err)
	}
	if err = SQLClient.Ping(); err != nil {
		log.Println("Warning: MySQL not connected. Ensure MySQL is running and database 'guvi_db' exists.", err)
	} else {
		fmt.Println("MySQL Connected")
	}
}
