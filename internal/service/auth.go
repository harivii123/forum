package service

import (
	"forum/internal/models"
	"time"
	"uuid"

	"golang.org/x/crypto/bcrypt"
)

// CheckPassword checks if the provided password matches the hashed password
func CheckPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

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
