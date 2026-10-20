package handlers

import (
	"errors"
	"log"
	"net/http"

	"forum/internal/database"
)

func (h *Holder) Logout(w http.ResponseWriter, r *http.Request) {
	log.Println("Logout handler called")
	if r.Method != http.MethodPost {
		h.err.Er, h.err.Code = errors.New("Method not allowed"), http.StatusMethodNotAllowed
		h.ErrorOut(w)
		return
	}

	cookie, err := r.Cookie("session_token")
	if err == nil {
		if err := database.DeleteSession(h.db, cookie.Value); err != nil {
			log.Printf("DeleteSession failed: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		log.Println("Session deleted from database")
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
