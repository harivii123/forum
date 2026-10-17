package handlers

import "net/http"

func clearSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
}

func (h *Holder) redirectIfLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	_, loggedIn, err := h.GetCurrentUser(w, r)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return false
	}

	if loggedIn {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return true
	}

	return false
}
