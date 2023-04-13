package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type MySQLDB struct {
	db *sql.DB
}

func New() *MySQLDB {
	//db, err := sql.Open("mysql", "root:mypassword@tcp(127.0.0.1:3306)/test")
	//db, err := sql.Open("mysql", "root:mypassword@tcp(127.0.0.1:3306)/test")

	db, err := sql.Open(`mysql`, `gameapp:gameappt0lk2o20@(127.0.0.1:3308)/gameapp_db`)
	if err != nil {
		panic(fmt.Errorf("can't open mysql %v", err))

	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySQLDB{db}
}
