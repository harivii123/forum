package handlers

import (
	"forum/internal/template"
	"io/fs"
	"net/http"
)

func Router(holder *Holder) *http.ServeMux {
	mux := http.NewServeMux()
	sub, _ := fs.Sub(template.Statics, "static")
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(sub))))
	mux.HandleFunc("/{$}", holder.LoadFrontPage)
	mux.HandleFunc("POST /post", holder.CreatePost)
	mux.HandleFunc("/login", holder.LoadRegistryPage)
	mux.HandleFunc("/profile", holder.LoadProfilePage)
	mux.HandleFunc("POST /posts/{id}/vote", holder.VoteOnPost)
	mux.HandleFunc("POST /{ID}/comment", holder.AddComment)

	return mux
}
