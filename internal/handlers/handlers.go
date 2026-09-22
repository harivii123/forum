package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"forum/internal/database"
	"forum/internal/models"
	"forum/internal/service"
	"log"
	"net/http"
	"time"
)

func (h *Holder) LoadRegistryPage(w http.ResponseWriter, r *http.Request) {
	// GET /login → show login component
	if r.Method == http.MethodGet {
		loginPage := h.engine.Render("index.html", map[string]any{
			"ShowLogin": true,
		})
		w.WriteHeader(http.StatusOK)
		w.Write(loginPage)
		return
	}
	// POST /login → process login
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")
		// 1. Find user
		user, err := database.GetUserByEmail(h.db, email)
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// 2. Check password
		err = service.CheckPassword(user.PasswordHash, password)
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		// 3. Create UUID session
		session := service.NewSession(user.ID)

		// 4. Save session to DB
		err = database.CreateSession(h.db, session)
		if err != nil {
			log.Printf("Error creating session: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// 5. Give session token to browser
		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    session.Token,
			Path:     "/",
			Expires:  session.ExpiresAt,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		})
		// 6. Login finished / Redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func (h *Holder) LoadFrontPage(w http.ResponseWriter, r *http.Request) {
	user, loggedIn, err := h.GetCurrentUser(w, r)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	page := &models.FrontPage{}
	mainSearchValue := r.URL.Query().Get("search")
	log.Println(mainSearchValue)
	posts, err := service.MainSearch(mainSearchValue, h.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	page.Posts = posts
	frontPage := h.engine.Render("index.html", map[string]any{
		"FrontPage": page,
		"User":      user,
		"LoggedIn":  loggedIn,
	})
	w.WriteHeader(http.StatusOK)
	w.Write(frontPage)
	//something_here.Execute(w, match) //<-- dont know how to use the fancy template parser yet
}

//func (h *Holder) CreatePost(w http.ResponseWriter, r *http.Request) {
// http.Redirect(w, r, "/", http.StatusSeeOther)
//}

func (h *Holder) CreatePost(w http.ResponseWriter, r *http.Request) {
	var newPost models.Post

	// err := json.NewDecoder(r.Body).Decode(&newPost)
	// if err != nil {
	// 	fmt.Println(1)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	newPost.UserID, newPost.Title, newPost.Body, newPost.Created = 1, "hashBrowniies", "i Like Frogs", time.Now()
	err, httpStatus := service.CreatePost(newPost, h.db)
	if err != nil {
		http.Error(w, err.Error(), httpStatus)
		return
	}
	fmt.Println("ok")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Holder) LoadProfilePage(w http.ResponseWriter, r *http.Request) {
	// get stuff from db insert into page bang
	// can be maybe also used to checkout other profiles not just ur own?
}
