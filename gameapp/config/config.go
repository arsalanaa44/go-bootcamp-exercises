package config

import (
	"gameapp/repository/mysql"
	"gameapp/service/authservice"
)

type HTTPServer struct {
	Port int
}
type Config struct {
	Auth       authservice.Config
	HTTPServer HTTPServer
	Mysql      mysql.Config
}
