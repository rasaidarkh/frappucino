package main

import (
	"database/sql"
	"fmt"
	"frappuccino/internal/handlers"
	"frappuccino/internal/helpers"
	"frappuccino/pkg/config"
	"frappuccino/pkg/lib/logger"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println(helpers.CreateMd5Hash(""))
	cfg := config.LoadConfig()
	logger := logger.SetupPrettySlog(os.Stdout)

	db, err := sql.Open("postgres", cfg.MakeConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	httpSrv := handlers.NewAPIServer(
		"0.0.0.0:8080",
		db,
		logger,
		rdb,
	)
	httpSrv.Run()
}
