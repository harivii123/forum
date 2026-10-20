package main

import (
	"database/sql"
	"flag"
	"forum/internal/database"
	"forum/internal/handlers"
	"forum/internal/service"
	"forum/internal/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

const port = ":8080"

func main() {
	if err := os.MkdirAll("uploads", 0755); err != nil {
		log.Fatal(err)
	}
	seed := flag.Bool("seed", false, "fill an empty database with example data")
	flag.Parse()

	dbPath, err := filepath.Abs("./database.db")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database path:", dbPath)

	db, err := sql.Open("sqlite3", "./database.db?_foreign_keys=on")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	err = database.CreateSchemas(db)
	if err != nil {
		log.Fatal(err)
	}

	if *seed {
		if err := database.Seed(db); err != nil {
			log.Fatal(err)
		}
		log.Println("Example data seeded")
	}

	service.StartSessionCleanUp(db)

	templateEngine := template.NewEngine("")
	templateEngine.ParseTemplates()

	holder := handlers.NewHolder(db, templateEngine)

	router := handlers.Router(holder)

	server := &http.Server{
		Addr:    port,
		Handler: router,
		//Read timeout?
		// Write timeout?
	}

	log.Printf("Server running on http://localhost%s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
