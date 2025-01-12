package config

import (
	"fmt"
	"os"
)

type Config struct {
	Host          string `env:"DB_HOST"`
	User          string `env:"DB_USER"`
	Password      string `env:"DB_PASSWORD"`
	DBname        string `env:"DB_NAME"`
	Port          string `env:"DB_PORT"`
	RedisURI      string `env:"REDIS_URI"`
	RedisPassword string `env:"REDIS_PASSWORD"`
	RedisDB       int    `env:"REDIS_DB"`
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
