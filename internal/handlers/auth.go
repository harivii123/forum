package handlers

import (
	"database/sql"
	"errors"
	"forum/internal/models"
	"net/http"
	"time"
)

func (h *Holder) GetCurrentUser(w http.ResponseWriter, r *http.Request) (models.User, bool, error) {
	var user models.User
	cookie, err := r.Cookie("session_token")
	// guest
	if errors.Is(err, http.ErrNoCookie) {
		return user, false, nil
	}
	// Database/server error:
	if err != nil {
		return user, false, err
	}

	var userID int
	err = h.db.QueryRow(`SELECT user_id FROM session WHERE token = ? AND expires_at > ?`,
		cookie.Value, time.Now()).Scan(&userID)

	if errors.Is(err, sql.ErrNoRows) {
		return user, false, nil
	}

	if err != nil {
		return user, false, err
	}
	// update new Expired time in DB and Browser cookie
	err = h.db.QueryRow(`SELECT id, username, email FROM user WHERE id = ?`, userID).Scan(&user.ID, &user.Username, &user.Email)
	if errors.Is(err, sql.ErrNoRows) {
		return user, false, nil
	}

	if err != nil {
		return user, false, err
	}

	now := time.Now()
	newExpiry := now.Add(15 * time.Minute)
	_, err = h.db.Exec(`UPDATE session SET last_active = ?, expires_at = ? WHERE token = ?`,
		now,
		newExpiry,
		cookie.Value,
	)

	if err != nil {
		return user, false, err
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
