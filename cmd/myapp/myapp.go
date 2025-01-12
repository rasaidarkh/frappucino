package main

import (
	"context"
	"database/sql"
	"frappuccino/internal/handlers"
	"frappuccino/pkg/config"
	"log"
	"log/slog"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.ConfigLoad()
	db, err := sql.Open("postgres", cfg.MakeConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT 1;")
	if err != nil {
		log.Fatal(err)
	}
	var hello string
	err = stmt.QueryRow().Scan(&hello)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(hello)

	server := handlers.NewAPIServer("0.0.0.0:8000", http.NewServeMux(), db, &slog.Logger{}, context.Background())
	server.Run()
}
