package handlers

import (
	"fmt"
	"forum/internal/database"
	"html/template"
	"log"
	"net/http"
)

func (h *Holder) LoadRegistryPage(w http.ResponseWriter, r *http.Request) {
	// check stuff insert stuff bang
	// maybe login can be done here too. Let's see
}

func (h *Holder) LoadFrontPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./internal/templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	mainSearchValue := r.URL.Query().Get("search")
	log.Println(mainSearchValue)
	match, err := h.MainSearch(mainSearchValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(match)

	w.WriteHeader(http.StatusOK)
	err = t.Execute(w, match)
}

func (h *Holder) LoadPostPage(w http.ResponseWriter, r *http.Request) {
	//get sttuff from url/body
	// go to repo get data from there and build page with data
}

func (h *Holder) LoadProfilePage(w http.ResponseWriter, r *http.Request) {
	// get stuff from db insert into page bang
	// can be maybe also used to checkout other profiles not just ur own?
}

// takes value from main search bar and finds all matches using fts(fast text search)
func (h *Holder) MainSearch(mainSearchValue string) (match []database.SearchResult, err error) {

	if mainSearchValue == "" {
		return nil, nil
	}
	match, err = database.FindAMatch(h.db, mainSearchValue)
	if err != nil {
		return nil, nil
	}

	fmt.Println(match) //template.Execute(w, match)
	return match, nil
}
