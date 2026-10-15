package handlers

import (
	"net/http"
)

type Err struct {
	Er   error
	Code int
}

func (h *Holder) ErrorOut(w http.ResponseWriter) {
	// log.Println("here123")
	ErrorPage := h.engine.Render("error.html", "error", map[string]any{
		"Err":   h.err.Er,
		"Code":  h.err.Code,
		"Error": true,
	})
	w.WriteHeader(h.err.Code)
	w.Write(ErrorPage)
}
