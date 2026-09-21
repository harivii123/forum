package handlers

import (
	"fmt"
	"forum/internal/models"
	"forum/internal/service"
	"log"
	"net/http"
	"time"
)

func (h *Holder) LoadRegistryPage(w http.ResponseWriter, r *http.Request) {
	// check stuff insert stuff bang
	// maybe login can be done here too. Let's see
}

func (h *Holder) LoadFrontPage(w http.ResponseWriter, r *http.Request) {
	page := &models.FrontPage{}
	mainSearchValue := r.URL.Query().Get("search")
	log.Println(mainSearchValue)
	posts, err := service.MainSearch(mainSearchValue, h.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	page.Posts = posts
	frontPage := h.engine.Render("index.html", map[string]any{
		"FrontPage": page,
	})
	w.WriteHeader(http.StatusOK)
	w.Write(frontPage)
	//something_here.Execute(w, match) //<-- dont know how to use the fancy template parser yet
}

//func (h *Holder) CreatePost(w http.ResponseWriter, r *http.Request) {
// http.Redirect(w, r, "/", http.StatusSeeOther)
//}

func (h *Holder) CreatePost(w http.ResponseWriter, r *http.Request) {
	var newPost models.Post

	// err := json.NewDecoder(r.Body).Decode(&newPost)
	// if err != nil {
	// 	fmt.Println(1)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	newPost.UserID, newPost.Title, newPost.Body, newPost.Created = 1, "hashBrowniies", "i Like Frogs", time.Now()
	err, httpStatus := service.CreatePost(newPost, h.db)
	if err != nil {
		http.Error(w, err.Error(), httpStatus)
		return
	}
	fmt.Println("ok")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Holder) LoadProfilePage(w http.ResponseWriter, r *http.Request) {
	// get stuff from db insert into page bang
	// can be maybe also used to checkout other profiles not just ur own?
}
