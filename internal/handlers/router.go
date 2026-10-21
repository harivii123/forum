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
	mux.Handle("GET /uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))
	mux.HandleFunc("/{$}", holder.LoadFrontPage)
	mux.HandleFunc("POST /post", holder.CreatePost)
	mux.HandleFunc("/login", holder.LoadRegistryPage)
	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			holder.LoadRegistrationPage(w, r)
			return
		}

		if r.Method == http.MethodPost {
			holder.Register(w, r)
			return
		}

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})
	mux.HandleFunc("/logout", holder.Logout)
	// mux.HandleFunc("/profile", holder.LoadProfilePage)
	mux.HandleFunc("GET /users/{id}", holder.LoadPublicProfilePage)
	mux.HandleFunc("POST /profile/edit", holder.UpdateProfile)
	mux.HandleFunc("POST /posts/{id}/vote", holder.VoteOnPost)
	mux.HandleFunc("POST /{ID}/comment", holder.AddComment)
	mux.HandleFunc("POST /comments/{id}/vote", holder.VoteOnComment)

	return mux
}
