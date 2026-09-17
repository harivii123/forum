package handlers

import "net/http"

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

// func (h *Holder) ToBeContinuedMaybe(?){
// } 
