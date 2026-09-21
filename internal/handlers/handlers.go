package handlers

import (
	"database/sql"
	"errors"
	"forum/internal/models"
	"forum/internal/service"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
		var user models.User

		row := h.db.QueryRow(`SELECT id, username, email, password_hash FROM user WHERE email = ?`, email)

		err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		// 2. Check password
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		// 3. Create UUID session
		var session models.Session
		session.Token = uuid.New().String()
		session.UserID = user.ID

		now := time.Now()
		session.CreatedAt = now
		session.LastActive = now
		session.ExpiresAt = now.Add(15 * time.Minute)

		// 4. Save session to DB
		_, err = h.db.Exec(`INSERT INTO session (token, user_id, created_at, expires_at, last_active) VALUES (?, ?, ?, ?, ?)`, session.Token, session.UserID, session.CreatedAt, session.ExpiresAt, session.LastActive)
		if err != nil {
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

func (h *Holder) LoadPostPage(w http.ResponseWriter, r *http.Request) {
	//get sttuff from url/body
	// go to repo get data from there and build page with data
}

func (h *Holder) LoadProfilePage(w http.ResponseWriter, r *http.Request) {
	// get stuff from db insert into page bang
	// can be maybe also used to checkout other profiles not just ur own?
}
