package handlers

import "net/http"

func (h *Holder) LoadRegistryPage(w http.ResponseWriter, r *http.Request) {
	// check stuff insert stuff bang
	// maybe login can be done here too. Let's see
}

func (h *Holder) LoadFrontPage(w http.ResponseWriter, r *http.Request) {
		//parse error html
		// status no bueno 
		// return
	
	//parse html
	// get css
	// get info from db
	//build page
	// query handling 
	// formvalue 
	// all good stuff
	// 
}

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
