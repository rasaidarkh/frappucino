package config

import (
	"fmt"
	"os"
)

type Config struct {
	ConnectionString string
}

func ConfigLoad() Config {
	var Cfg Config
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	Cfg.ConnectionString = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
		host, port, user, password, dbname,
	)
	fmt.Println(Cfg.ConnectionString)
	return Cfg
}
