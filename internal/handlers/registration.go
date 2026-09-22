package handlers

import (
	"database/sql"
	"errors"
	"forum/internal/database"
	"forum/internal/models"
	"forum/internal/service"
	"net/http"
)

func (h *Holder) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	user.Username = r.FormValue("username")
	user.Email = r.FormValue("email")
	password := r.FormValue("password")

	_, err := database.GetUserByUsername(h.db, user.Username)
	if err == nil {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}
	if !errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	_, err = database.GetUserByEmail(h.db, user.Email)
	if err == nil {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}
	if !errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	passwordHash, err := service.HashPassword(password)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	user.PasswordHash = passwordHash

	user.ID, err = database.CreateUser(h.db, user)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	session := service.NewSession(user.ID)
	err = database.CreateSession(h.db, session)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    session.Token,
		Path:     "/",
		Expires:  session.ExpiresAt,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
