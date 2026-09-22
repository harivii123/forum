package handlers

import (
	"database/sql"
	"errors"
	"forum/internal/database"
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

	userID, err := database.GetSessionUserID(h.db, cookie.Value)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, false, nil
	}

	if err != nil {
		return nil, false, err
	}

	// update new Expired time in DB and Browser cookie
	user, err := database.GetUserByID(h.db, userID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, false, nil
	}

	if err != nil {
		return nil, false, err
	}

	now, newExpiry := service.RefreshSessionTime()
	err = database.UpdateSessionActivity(h.db, cookie.Value, now, newExpiry)

	if err != nil {
		return nil, false, err
	}

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
