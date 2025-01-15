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
	Cfg.Host = getEnv("DB_HOST", "localhost")
	Cfg.User = getEnv("DB_USER", "postgres")
	Cfg.Password = getEnv("DB_PASSWORD", "0000")
	Cfg.DBname = getEnv("DB_NAME", "frappuccino_db")
	Cfg.Port = getEnv("DB_PORT", "5432")

	return Cfg
}

func (c *Config) MakeConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DBname,
	)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
