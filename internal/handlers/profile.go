package handlers

import (
	"database/sql"
	"errors"
	"forum/internal/database"
	"forum/internal/models"
	"forum/internal/service"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
)

// UpdateProfile saves the bio of the currently logged-in user.
func (h *Holder) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.err.Er = errors.New("Method not allowed")
		h.err.Code = http.StatusMethodNotAllowed
		h.ErrorOut(w)
		return
	}

	user, loggedIn, err := h.GetCurrentUser(w, r)
	if err != nil {
		h.err.Er = err
		h.err.Code = http.StatusInternalServerError
		h.ErrorOut(w)
		return
	}

	if !loggedIn {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	bio := strings.TrimSpace(r.FormValue("bio"))
	if utf8.RuneCountInString(bio) > 500 {
		h.err.Er = errors.New("Bio must be 500 characters or fewer")
		h.err.Code = http.StatusBadRequest
		h.ErrorOut(w)
		return
	}

	err = database.UpdateUserProfile(h.db, user.ID, bio)
	if err != nil {
		h.err.Er = err
		h.err.Code = http.StatusInternalServerError
		h.ErrorOut(w)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Holder) LoadPublicProfilePage(w http.ResponseWriter, r *http.Request) {
	viewer, loggedIn, err := h.GetCurrentUser(w, r)
	if err != nil {
		h.err.Er, h.err.Code = err, http.StatusInternalServerError
		h.ErrorOut(w)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		http.NotFound(w, r)
		return
	}

	user, err := database.GetUserByID(h.db, id)
	if errors.Is(err, sql.ErrNoRows) {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		h.ErrorOut(w)
		return
	}
	var profile models.UserView
	if viewer != nil {
		profile.ID = user.ID
		profile.Username = user.Username
		profile.ProfilePicture = user.ProfilePicture
	}

	data := map[string]any{
		"LoggedIn":        loggedIn,
		"Profile":         profile,
		"Owner":           user,
		"User":            viewer,
		"IsPublicProfile": true,
		"IsOwnProfile":    loggedIn && viewer.ID == user.ID,
	}

	w.Write(h.engine.Render("index.html", "base", data))
}

func (h *Holder) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	user, loggedIn, err := h.GetCurrentUser(w, r)
	if err != nil {
		h.err.Er, h.err.Code = err, http.StatusInternalServerError
		h.ErrorOut(w)
		return
	}
	if !loggedIn {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	ctx := r.Context()
	if err := service.DeleteUser(ctx, h.db, user.ID); err != nil {
		h.err.Er, h.err.Code = err, http.StatusInternalServerError
		h.ErrorOut(w)
		return
	}

	clearSessionCookie(w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
