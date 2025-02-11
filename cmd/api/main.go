package main

import (
	"database/sql"
	"frappuccino/internal/handlers"
	"frappuccino/pkg/config"
	"frappuccino/pkg/lib/logger"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.LoadConfig()
	logger := logger.SetupPrettySlog(os.Stdout)

	db, _ := sql.Open("postgres", cfg.MakeConnectionString())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err = db.Ping(); err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	httpSrv := handlers.NewAPIServer(
		"0.0.0.0:8080",
		db,
		logger,
	)
	httpSrv.Run()
}
