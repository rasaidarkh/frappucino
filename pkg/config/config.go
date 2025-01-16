package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBHost        string `env:"DB_HOST"`
	DBUser        string `env:"DB_USER"`
	DBPassword    string `env:"DB_PASSWORD"`
	DBName        string `env:"DB_NAME"`
	DBPort        string `env:"DB_PORT"`
	JWTSecret     string `env:"JWT_SECRET"`
	RedisURI      string `env:"REDIS_URI"`
	RedisPassword string `env:"REDIS_PASSWORD"`
	RedisDB       int    `env:"REDIS_DB"`
}

var (
	Cfg Config
)

func LoadConfig() *Config {
	Cfg.DBHost = getEnv("DB_HOST", "localhost")
	Cfg.DBUser = getEnv("DB_USER", "postgres")
	Cfg.DBPassword = getEnv("DB_PASSWORD", "0000")
	Cfg.DBName = getEnv("DB_NAME", "frappuccino_db")
	Cfg.DBPort = getEnv("DB_PORT", "5432")
	Cfg.JWTSecret = getEnv("JWT_SECRET", "not-so-secret-now-is-it?")

	return &Cfg
}

func GetConfing() *Config {
	return &Cfg
}

func (c *Config) MakeConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName,
	)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
