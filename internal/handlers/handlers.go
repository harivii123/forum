package handlers

import (
	"forum/internal/service"
	"html/template"
	"net/http"
)

func (h *Holder) LoadRegistryPage(w http.ResponseWriter, r *http.Request) {
	// check stuff insert stuff bang
	// maybe login can be done here too. Let's see
}

func (h *Holder) LoadFrontPage(w http.ResponseWriter, r *http.Request) {
	pageData, err := h.GetFrontPageData()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	frontPage := h.engine.Render("index.html", map[string]any{
		"PageData": pageData,
	})

	w.Write(frontPage)
	t, err := template.ParseFiles("./internal/templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	mainSearchValue := r.URL.Query().Get("search")
	// log.Println(mainSearchValue)
	match, err := service.MainSearch(mainSearchValue, h.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// log.Println(match)

	w.WriteHeader(http.StatusOK)
	err = t.Execute(w, match)
}

//func (h *Holder) CreatePost(w http.ResponseWriter, r *http.Request) {
// http.Redirect(w, r, "/", http.StatusSeeOther)
//}

func (h *Holder) LoadPostPage(w http.ResponseWriter, r *http.Request) {
	//get sttuff from url/body
	// go to repo get data from there and build page with data
}

func (h *Holder) LoadProfilePage(w http.ResponseWriter, r *http.Request) {
	// get stuff from db insert into page bang
	// can be maybe also used to checkout other profiles not just ur own?
}

// takes value from main search bar and finds all matches using fts(fast text search)
