package handlers

import (
	"log"
	"net/http"
)

type Err struct {
	Er   error
	Code int
}

func (h *Holder) ErrorOut(w http.ResponseWriter) {
	log.Println("here")
	ErrorPage := h.engine.Render("index.html", map[string]any{
		"Err":  h.err.Er,
		"Code": h.err.Code,
	})
	w.WriteHeader(h.err.Code)
	w.Write(ErrorPage)
}
