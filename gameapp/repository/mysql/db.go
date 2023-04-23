package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Config struct {
	UserName string
	Password string
	Port     int
	Host     string
	DBName   string
}

type MySQLDB struct {
	config Config
	db     *sql.DB
}

func New(config Config) *MySQLDB {
	//db, err := sql.Open("mysql", "root:mypassword@tcp(127.0.0.1:3306)/test")
	//db, err := sql.Open("mysql", "root:mypassword@tcp(127.0.0.1:3306)/test")

	db, err := sql.Open(`mysql`,
		fmt.Sprintf(`%s:%s@(%s:%d)/%s`,
			config.UserName, config.Password, config.Host, config.Port, config.DBName,
		),
	)
	if err != nil {
		panic(fmt.Errorf("can't open mysql %v", err))

	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySQLDB{config, db}
}
