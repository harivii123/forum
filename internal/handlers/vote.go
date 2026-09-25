package handlers

import (
	"errors"
	"forum/internal/service"
	"net/http"
	"strconv"
)

func (h *Holder) VoteOnPost(w http.ResponseWriter, r *http.Request) {

	user, loggedIn, err := h.GetCurrentUser(w, r)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	if !loggedIn { //No voting without being logged in
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

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

	err = service.SetPostVote(h.db, user.ID, postID, vote)
	switch {
	case errors.Is(err, service.ErrPostNotFound):
		http.Error(w, "post not found", http.StatusNotFound)
		return
	case err != nil:
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
