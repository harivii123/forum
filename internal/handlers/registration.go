package handlers

import (
	"errors"
	"forum/internal/database"
	"forum/internal/service"
	"log"
	"net/http"
)

func (h *Holder) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := service.RegisterUser(h.db, username, email, password)
	if err != nil {
		// status := http.StatusBadRequest
		message := "Something went wrong. Please try again."

		switch {
		case errors.Is(err, service.ErrMissingFields):
			message = "Please fill in all fields"
		case errors.Is(err, service.ErrInvalidEmail):
			message = "Please enter a valid email address"
		case errors.Is(err, service.ErrUsernameTaken):
			// status = http.StatusConflict
			message = "This username is already taken"
		case errors.Is(err, service.ErrEmailTaken):
			// status = http.StatusConflict
			message = "This email is already registered"
		default:
			// status = http.StatusInternalServerError
			log.Printf("Registration failed: %v", err)
		}

		h.RenderFrontPage(w, r, "register", message)
		return
	}

	session := service.NewSession(user.ID)
	err = database.CreateSession(h.db, session)
	if err != nil {
		log.Printf("Error creating session after registration: %v", err)
		h.RenderFrontPage(w, r, "register", "Your account was created, but we couldn't log you in. Please try logging in.")
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
