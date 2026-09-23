package main

import (
	"database/sql"
	"flag"
	"forum/internal/database"
	"forum/internal/handlers"
	"forum/internal/template"
	"io/fs"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

const port = ":8080"

func main() {
	seed := flag.Bool("seed", false, "fill an empty database with example data")
	flag.Parse()

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
	//log.Printf("Tables:%s", tables)

	templateEngine := template.NewEngine("")
	templateEngine.ParseTemplates()

	holder := handlers.NewHolder(db, templateEngine)

	mux := http.NewServeMux()
	sub, _ := fs.Sub(template.Statics, "static")
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(sub))))
	mux.HandleFunc("/{$}", holder.LoadFrontPage)
	mux.HandleFunc("POST /post", holder.CreatePost)
	mux.HandleFunc("/login", holder.LoadRegistryPage)
	mux.HandleFunc("/profile", holder.LoadProfilePage)
	mux.HandleFunc("POST /{ID}/comment", holder.AddComment)

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
