package main

import (
	"database/sql"
	"html/template"
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

	t, err := template.ParseFiles("") //html paths here
	if err != nil {
		log.Fatal(err)
	}
	
	fs := http.FileServer(http.Dir("./static/"))

	mux := http.NewServeMux()

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
