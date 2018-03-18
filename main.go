package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"

	_ "github.com/lib/pq"
)

func main() {
	pgConnStr := os.Getenv("PGCONN")
	if pgConnStr == "" {
		log.Print("PGCONN environment variable required")
		os.Exit(1)
	}
	dbo, err := sql.Open("postgres", os.Getenv("PGCONN"))
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	err = dbo.Ping()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	query := os.Args[1]
	cmd := os.Args[2]
	executeQuery(query, dbo)
}

func executeQuery(query string, dbo *sql.DB) {
	_, err := dbo.Exec(query)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
