package handlers

import (
	"errors"
	"forum/internal/models"
	"forum/internal/service"
	"net/http"
)

func (h *Holder) GetCurrentUser(w http.ResponseWriter, r *http.Request) (*models.User, bool, error) {
	cookie, err := r.Cookie("session_token")
	// guest
	if errors.Is(err, http.ErrNoCookie) {
		return nil, false, nil
	}
	// Database/server error:
	if err != nil {
		return nil, false, err
	}

	// check if session is active and valid, else clear cookie and treat as guest
	user, newExpiry, err := service.GetAuthenticatedUser(h.db, cookie.Value)
	if errors.Is(err, service.ErrInvalidSession) {
		clearSessionCookie(w)
		return nil, false, nil
	}

	if err != nil {
		return nil, false, err
	}

	// Refresh the browser cookie expiration time
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    cookie.Value,
		Path:     "/",
		Expires:  newExpiry,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	return user, true, nil
}
