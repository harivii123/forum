package service

import (
	"database/sql"
	"errors"
	"forum/internal/database"
	"forum/internal/models"
	"log"
	"time"
	"uuid"
)

// Create UUID session
func NewSession(userID int) models.Session {
	var session models.Session
	session.Token = uuid.New().String()
	session.UserID = userID
	now, expiry := RefreshSessionTime()
	session.CreatedAt = now
	session.LastActive = now
	session.ExpiresAt = expiry
	return session
}

// RefreshSessionTime refreshes the session time and returns the new last active time and new expiry time
func RefreshSessionTime() (time.Time, time.Time) {
	now := time.Now()
	newExpiry := now.Add(15 * time.Minute)
	return now, newExpiry
}

// GetAuthenticatedUser validates a session, retrieves the associated user,
// and refreshes the session expiration time
func GetAuthenticatedUser(db *sql.DB, token string) (*models.User, time.Time, error) {
	// check if session is active and valid, else delete session and treat as guest
	userID, err := database.GetSessionUserID(db, token)
	if errors.Is(err, sql.ErrNoRows) {
		if deleteErr := database.DeleteSession(db, token); deleteErr != nil {
			return nil, time.Time{}, deleteErr
		}
		return nil, time.Time{}, ErrInvalidSession
	}

	if err != nil {
		return nil, time.Time{}, err
	}

	user, err := database.GetUserByID(db, userID)
	if errors.Is(err, sql.ErrNoRows) {
		if deleteErr := database.DeleteSession(db, token); deleteErr != nil {
			return nil, time.Time{}, deleteErr
		}
		return nil, time.Time{}, ErrInvalidSession
	}

	if err != nil {
		return nil, time.Time{}, err
	}

	now, newExpiry := RefreshSessionTime()
	err = database.UpdateSessionActivity(db, token, now, newExpiry)
	// The session may have expired before the activity update
	if errors.Is(err, sql.ErrNoRows) {
		if deleteErr := database.DeleteSession(db, token); deleteErr != nil {
			return nil, time.Time{}, deleteErr
		}
		return nil, time.Time{}, ErrInvalidSession
	}
	if err != nil {
		return nil, time.Time{}, err
	}

	return user, newExpiry, nil
}

// It periodically deletes expired sessions
func StartSessionCleanUp(db *sql.DB) {
	ticker := time.NewTicker(15 * time.Minute)

	go func() {
		defer ticker.Stop()

		for range ticker.C {
			if err := database.DeleteExpiredSessions(db, time.Now()); err != nil {
				log.Printf("Failed to clean up expired sessions: %v", err)
			}
		}
	}()
}
