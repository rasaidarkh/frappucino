package config

import (
	"fmt"
	"os"
)

type Config struct {
	Host     string
	User     string
	Password string
	DBname   string
	Port     string
}

func ConfigLoad() Config {
	var Cfg Config
	Cfg.Host = os.Getenv("DB_HOST")
	Cfg.User = os.Getenv("DB_USER")
	Cfg.Password = os.Getenv("DB_PASSWORD")
	Cfg.DBname = os.Getenv("DB_NAME")
	Cfg.Port = os.Getenv("DB_PORT")

	return Cfg
}

func (c *Config) MakeConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DBname,
	)
}
