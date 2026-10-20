package service

import (
	"database/sql"
	"errors"

	"forum/internal/database"
	"forum/internal/models"

	"time"
)

// GetAuthenticatedUser validates a session, retrieves the associated user,
// and refreshes the session expiration time
func GetAuthenticatedUser(db *sql.DB, token string) (*models.User, time.Time, error) {
	// check if session is active and valid, else treat as guest
	userID, err := database.GetSessionUserID(db, token)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, time.Time{}, ErrInvalidSession
	}

	if err != nil {
		return nil, time.Time{}, err
	}

	user, err := database.GetUserByID(db, userID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, time.Time{}, ErrInvalidSession
	}

	if err != nil {
		return nil, time.Time{}, err
	}

	now, newExpiry := RefreshSessionTime()
	err = database.UpdateSessionActivity(db, token, now, newExpiry)
	// The session may have expired before the activity update
	if errors.Is(err, sql.ErrNoRows) {
		return nil, time.Time{}, ErrInvalidSession
	}
	if err != nil {
		return nil, time.Time{}, err
	}

	return user, newExpiry, nil
}
