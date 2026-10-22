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
		h.err.Er, h.err.Code = err, http.StatusInternalServerError
		h.ErrorOut(w)
		return
	}
	if !loggedIn {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	postID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		h.err.Er, h.err.Code = err, http.StatusBadRequest
		h.ErrorOut(w)
		return
	}

	vote, err := strconv.Atoi(r.FormValue("vote"))
	if err != nil {
		h.err.Er, h.err.Code = err, http.StatusBadRequest
		h.ErrorOut(w)
		return
	}

	err = service.SetPostVote(h.db, user.ID, postID, vote)
	switch {
	case errors.Is(err, service.ErrPostNotFound):
		h.err.Er, h.err.Code = errors.New("Post not found"), http.StatusNotFound
		h.ErrorOut(w)
		return
	case errors.Is(err, service.ErrInvalidVote):
		h.err.Er, h.err.Code = errors.New("Invalid vote"), http.StatusBadRequest
		h.ErrorOut(w)
		return
	case err != nil:
		h.err.Er, h.err.Code = err, http.StatusInternalServerError
		h.ErrorOut(w)
		return
	}
	query := r.FormValue("redirect")

	http.Redirect(w, r, "/?"+query+"#post-"+strconv.Itoa(postID), http.StatusSeeOther)
}

func (h *Holder) VoteOnComment(w http.ResponseWriter, r *http.Request) {
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

	commentID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		h.err.Er, h.err.Code = errors.New("Invalid comment id"), http.StatusBadRequest
		h.ErrorOut(w)
		return
	}

	vote, err := strconv.Atoi(r.FormValue("vote"))
	if err != nil {
		h.err.Er, h.err.Code = errors.New("Invalid vote"), http.StatusBadRequest
		h.ErrorOut(w)
		return
	}

	err = service.SetCommentVote(h.db, user.ID, commentID, vote)
	switch {
	case errors.Is(err, service.ErrCommentNotFound):
		h.err.Er, h.err.Code = errors.New("Comment not found"), http.StatusNotFound
		h.ErrorOut(w)
		return
	case errors.Is(err, service.ErrInvalidVote):
		h.err.Er, h.err.Code = errors.New("Invalid vote"), http.StatusBadRequest
		h.ErrorOut(w)
		return
	case err != nil:
		h.err.Er, h.err.Code = err, http.StatusInternalServerError
		h.ErrorOut(w)
		return
	}

	query := r.FormValue("redirect")
	http.Redirect(w, r, "/?"+query, http.StatusSeeOther)
}
