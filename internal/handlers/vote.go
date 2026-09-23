package handlers

import (
	"log"
	"net/http"
	"strconv"
)

func (h *Holder) VoteOnPost(w http.ResponseWriter, r *http.Request) {
	postID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "bad post id", http.StatusBadRequest)
		return
	}
	vote, err := strconv.Atoi(r.FormValue("vote"))
	if err != nil {
		http.Error(w, "bad vote", http.StatusBadRequest)
		return
	}
	user, loggedIn, err := h.GetCurrentUser(w, r)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	if !loggedIn { //I think no voting without being logged in?
		http.Error(w, "cannot vote without logging in", http.StatusForbidden)
		return
	}

	log.Printf("user %d voted %d on post %d", user.ID, vote, postID)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
