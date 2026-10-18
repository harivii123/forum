package handlers

import (
	"database/sql"
	"errors"
	"forum/internal/database"
	"forum/internal/models"
	"forum/internal/service"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"uuid"
)

func (h *Holder) LoadRegistryPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet && h.redirectIfLoggedIn(w, r) {
		return
	}
	// GET /login → show login component
	if r.Method == http.MethodGet {
		if h.redirectIfLoggedIn(w, r) {
			return
		}

		h.RenderFrontPage(w, r, "login")
		return
	}
	// POST /login → process login
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")
		// 1. Find user
		user, err := database.GetUserByEmail(h.db, email)
		if errors.Is(err, sql.ErrNoRows) {
			h.err.Er, h.err.Code = errors.New("Invalid email or password"), http.StatusUnauthorized
			h.ErrorOut(w)
			return
		}

		if err != nil {
			h.err.Er, h.err.Code = err, http.StatusInternalServerError
			h.ErrorOut(w)
			return
		}

		// 2. Check password
		err = service.CheckPassword(user.PasswordHash, password)
		if err != nil {
			h.err.Er, h.err.Code = errors.New("Invalid email or password"), http.StatusUnauthorized
			h.ErrorOut(w)
			return
		}

		// 3. Create UUID session
		session := service.NewSession(user.ID)

		// 4. Save session to DB
		err = database.CreateSession(h.db, session)
		if err != nil {
			log.Printf("Error creating session: %v", err)
			h.err.Er, h.err.Code = err, http.StatusInternalServerError
			h.ErrorOut(w)
			return
		}

		// 5. Give session token to browser
		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    session.Token,
			Path:     "/",
			Expires:  session.ExpiresAt,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		})
		// 6. Login finished / Redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func (h *Holder) RenderFrontPage(w http.ResponseWriter, r *http.Request, showAuth string) {
	var f models.Filters
	f.PageFilters(r.URL.Query())
	user, loggedIn, err := h.GetCurrentUser(w, r)
	if err != nil {
		h.err.Er, h.err.Code = err, http.StatusInternalServerError
		h.ErrorOut(w)
		return
	}
	var profile models.UserView
	if user != nil {
		profile.ID = user.ID
		profile.Username = user.Username
		profile.ProfilePicture = ""
	}

	var userID int
	if user != nil {
		userID = user.ID
	}

	filters, args := filtersFromQuery(r.URL.Query(), userID, loggedIn)
	ctx := r.Context()

	posts, err := service.FilterPosts(filters, args, ctx, h.db)
	if err != nil {
		h.err.Er, h.err.Code = err, http.StatusInternalServerError
		h.ErrorOut(w)
		return
	}

	var categories []models.CategoryView

	categories, err = service.GetCategories(ctx, h.db)

	// post_error is a one-time message from CreatePost; drop it from Query so
	// vote redirects don't carry it along and show it again
	q := r.URL.Query()
	postErr := postErrorMessages[q.Get("post_error")]
	q.Del("post_error")
	query := q.Encode()

	frontPage := h.engine.Render("index.html", "base", map[string]any{
		"Posts":        posts,
		"Profile":      profile,
		"Categories":   categories,
		"User":         user,
		"LoggedIn":     loggedIn,
		"Filters":      f,
		"Query":        query,
		"ShowLogin":    showAuth == "login",
		"ShowRegister": showAuth == "register",
		"PostError":    postErr,
	})
	w.WriteHeader(http.StatusOK)
	w.Write(frontPage)
	//something_here.Execute(w, match) //<-- dont know how to use the fancy template parser yet
}

func (h *Holder) LoadFrontPage(w http.ResponseWriter, r *http.Request) {
	h.RenderFrontPage(w, r, "")
}

func (h *Holder) LoadRegistrationPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet && h.redirectIfLoggedIn(w, r) {
		return
	}

	if r.Method != http.MethodGet {
		h.err.Er, h.err.Code = errors.New("Method not allowed"), http.StatusMethodNotAllowed
		h.ErrorOut(w)
		return
	}

	h.RenderFrontPage(w, r, "register")
}

func (h *Holder) CreatePost(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 20<<20)

	//Check if user is logged in
	_, ok, err := h.GetCurrentUser(w, r)
	if err != nil {
		h.err.Er, h.err.Code = err, http.StatusInternalServerError
		h.ErrorOut(w)
		return
	}
	if !ok {
		io.Copy(io.Discard, r.Body)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if err := r.ParseMultipartForm(20 << 20); err != nil {
		var tooBig *http.MaxBytesError
		if errors.As(err, &tooBig) {
			postError(w, r, "too_large")
			return
		}
		h.err.Er, h.err.Code = err, http.StatusBadRequest
		h.ErrorOut(w)
		return
	}
	if strings.TrimSpace(r.FormValue("title")) == "" || strings.TrimSpace(r.FormValue("body")) == "" {
		postError(w, r, "empty")
		return
	}
	file, header, err := r.FormFile("image")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		h.err.Er, h.err.Code = err, http.StatusInternalServerError
		h.ErrorOut(w)
		return
	}
	var imagePath string
	var path string
	if file != nil {
		defer file.Close()
		if header.Size > 5<<20 {
			postError(w, r, "too_large")
			return
		}
		isImage := make([]byte, 512)
		n, err := file.Read(isImage)
		if err != nil {
			h.err.Er, h.err.Code = err, http.StatusBadRequest
			h.ErrorOut(w)
			return
		}
		fileType := http.DetectContentType(isImage[:n])
		name := uuid.New().String()
		var end string
		switch fileType {
		case "image/png":
			end = ".png"
		case "image/jpeg":
			end = ".jpeg"
		case "image/gif":
			end = ".gif"
		default:
			postError(w, r, "bad_type")
			return
		}
		_, err = file.Seek(0, 0)
		if err != nil {
			h.err.Er, h.err.Code = err, http.StatusInternalServerError
			h.ErrorOut(w)
			return
		}
		imagePath = "/uploads/" + name + end
		path = filepath.Join("uploads", name+end)
		dst, err := os.Create(path)
		if err != nil {
			h.err.Er, h.err.Code = err, http.StatusInternalServerError
			h.ErrorOut(w)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			h.err.Er, h.err.Code = err, http.StatusInternalServerError
			os.Remove(path)
			h.ErrorOut(w)
			return
		}
	}

	newPost := models.Post{
		Title: r.FormValue("title"),
		Body:  r.FormValue("body"),
		Image: imagePath,
	}
	//user data from request cookies
	cookie, err := r.Cookie("session_token")
	//createpost returns err and appropriate http code
	err, httpStatus := service.CreatePost(cookie, newPost, h.db)
	if err != nil {
		h.err.Er, h.err.Code = err, http.StatusInternalServerError
		if path != "" {
			os.Remove(path)
		}
		h.ErrorOut(w)
		return
	}

	//redirect back to homepage if everything went well
	http.Redirect(w, r, "/", httpStatus)
}

func (h *Holder) LoadProfilePage(w http.ResponseWriter, r *http.Request) {
	// get stuff from db insert into page bang
	// can be maybe also used to checkout other profiles not just ur own?
}

func (h *Holder) AddComment(w http.ResponseWriter, r *http.Request) {
	//get the post_id from request
	_, loggedIn, err := h.GetCurrentUser(w, r)
	if err != nil {
		h.err.Er, h.err.Code = err, http.StatusInternalServerError
		h.ErrorOut(w)
		return
	}
	if !loggedIn {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	postIDStr := r.PathValue("ID")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		h.err.Er, h.err.Code = err, http.StatusInternalServerError
		h.ErrorOut(w)
		return
	}
	//get comment value from request
	newComment := models.Comment{
		Post_id: postID,
		Body:    r.FormValue("comment"),
	}
	//get user data from request cookies
	cookie, err := r.Cookie("session_token")
	if err != nil {
		h.err.Er, h.err.Code = err, http.StatusInternalServerError
		h.ErrorOut(w)
		return
	}
	//service.CreateComment returns error and appropriate httpstatus
	err, httpStatus := service.CreateComment(newComment, cookie, h.db)
	if err != nil {
		h.err.Er, h.err.Code = err, httpStatus
		h.ErrorOut(w)
		return
	}
	//redirect back to homepage if everything went well
	http.Redirect(w, r, "/", httpStatus)
}

func filtersFromQuery(q url.Values, userID int, loggedIn bool) (string, []any) {
	filters := []string{}
	args := []any{}

	if search := q.Get("search"); search != "" {
		param, arg := service.Search(search)
		filters = append(filters, param)
		args = append(args, arg)
	}

	if posts := q.Get("posts"); posts == "Posts" && loggedIn {
		param, arg := service.Posts(userID)
		filters = append(filters, param)
		args = append(args, arg)
	}

	if likes := q.Get("likes"); likes == "Likes" && loggedIn {
		param, arg := service.Likes(userID)
		filters = append(filters, param)
		args = append(args, arg)
	}

	var ids []int
	for _, id := range q["category"] {
		if id, err := strconv.Atoi(id); err == nil {
			ids = append(ids, id)
		}
	}
	if len(ids) > 0 {
		param, arg := service.Categories(ids)
		filters = append(filters, param)
		args = append(args, arg...)
	}
	var result string
	if len(filters) > 0 {
		result = " WHERE " + strings.Join(filters, " AND ")
	}

	if order := q.Get("order"); order != "" {
		param := service.Order(order)
		result += " GROUP BY p.id "
		result += param + ";"
	} else {
		result += " GROUP BY p.id ORDER BY p.created_at DESC;"
	}

	return result, args
}

// user-facing messages for create post validation errors, keyed by the
// post_error query value so arbitrary text can't be injected via the URL
var postErrorMessages = map[string]string{
	"empty":     "Title and text cannot be empty.",
	"too_large": "Image is too large (max 5 MB).",
	"bad_type":  "Only PNG, JPEG and GIF images are allowed.",
}

// sends the user back to the front page where the form shows the error
func postError(w http.ResponseWriter, r *http.Request, code string) {
	http.Redirect(w, r, "/?post_error="+code, http.StatusSeeOther)
}
