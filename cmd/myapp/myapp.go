package myapp

import (
	"database/sql"
	"frappuccino/internal/config"
	"log"
)

func main() {
	cfg := config.ConfigLoad()
	db, err := sql.Open("postgres", cfg.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT 'HALOO'")
	if err != nil {
		log.Fatal(err)
	}
	var hello string
	err = stmt.QueryRow().Scan(&hello)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(hello)
}
