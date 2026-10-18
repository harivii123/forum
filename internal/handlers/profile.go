package handlers

import (
	"errors"
	"forum/internal/database"
	"net/http"
	"strings"
	"unicode/utf8"
)

// LoadEditProfilePage displays the profile editing form for the currently logged-in user.
// func (h *Holder) LoadEditProfilePage(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodGet {
	// 	h.err.Er = errors.New("Method not allowed")
	// 	h.err.Code = http.StatusMethodNotAllowed
	// 	h.ErrorOut(w)
	// 	return
	// }

// 	user, loggedIn, err := h.GetCurrentUser(w, r)
// 	if err != nil {
// 		h.err.Er = err
// 		h.err.Code = http.StatusInternalServerError
// 		h.ErrorOut(w)
// 		return
// 	}

// 	if !loggedIn {
// 		http.Redirect(w, r, "/login", http.StatusSeeOther)
// 		return
// 	}

// 	page := h.engine.Render("profile-edit.html", "profile-edit", map[string]any{
// 		"User": user,
// 	})

// 	w.WriteHeader(http.StatusOK)
// 	_, _ = w.Write(page)

// }

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
