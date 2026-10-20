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

	if r.Method != http.MethodPost {
		h.err.Er, h.err.Code = errors.New("Method not allowed"), http.StatusMethodNotAllowed
		h.ErrorOut(w)
		return
	}

	user.Username = r.FormValue("username")
	user.Email = r.FormValue("email")
	password := r.FormValue("password")

	if user.Username == "" || user.Email == "" || password == "" {
		h.err.Er, h.err.Code = errors.New("Username, email and password are required"), http.StatusBadRequest
		h.ErrorOut(w)
		return
	}

	_, err := database.GetUserByUsername(h.db, user.Username)
	if err == nil {
		h.err.Er, h.err.Code = errors.New("Username already exists"), http.StatusConflict
		h.ErrorOut(w)
		return
	}
	if !errors.Is(err, sql.ErrNoRows) {
		h.err.Er, h.err.Code = err, http.StatusInternalServerError
		h.ErrorOut(w)
		return
	}

	_, err = database.GetUserByEmail(h.db, user.Email)
	if err == nil {
		h.err.Er, h.err.Code = errors.New("Email already exists"), http.StatusConflict
		h.ErrorOut(w)
		return
	}
	if !errors.Is(err, sql.ErrNoRows) {
		h.err.Er, h.err.Code = err, http.StatusInternalServerError
		h.ErrorOut(w)
		return
	}

	passwordHash, err := service.HashPassword(password)
	if err != nil {
		h.err.Er, h.err.Code = err, http.StatusInternalServerError
		h.ErrorOut(w)
		return
	}
	user.PasswordHash = passwordHash

	user.ID, err = database.CreateUser(h.db, user)

	if err != nil {
		h.err.Er, h.err.Code = err, http.StatusInternalServerError
		h.ErrorOut(w)
		return
	}

	session := service.NewSession(user.ID)
	err = database.CreateSession(h.db, session)
	if err != nil {
		h.err.Er, h.err.Code = err, http.StatusInternalServerError
		h.ErrorOut(w)
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
