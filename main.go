package main

import (
	"database/sql"
	"forum/internal/database"
	"forum/internal/handlers"
	"forum/internal/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

const port = ":8080"

func main() {

	db, err := sql.Open("sqlite3", "./database.db?_foreign_keys=on")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	tables, err := database.CreateSchema(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Tables:%s", tables)

	templateEngine := template.NewEngine("")
	templateEngine.ParseTemplates()

	holder := handlers.NewHolder(db, templateEngine)

	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", holder.LoadFrontPage)
	mux.HandleFunc("/post", holder.LoadPostPage)
	mux.HandleFunc("/login", holder.LoadRegistryPage)
	mux.HandleFunc("/profile", holder.LoadProfilePage)

	//handlefunc yada yada

	server := &http.Server{
		Addr:    port,
		Handler: mux,
		//Read timeout?
		// Write timeout?
	}

	log.Printf("Server running on http://localhost%s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
