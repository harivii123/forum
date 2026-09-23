package service

import (
	"context"
	"database/sql"
	"errors"
	"forum/internal/database"
	"forum/internal/models"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// takes value from main search bar and finds all matches using fts(fast text search)
func FilterPosts(filters string, args []any, ctx context.Context, db *sql.DB) ([]models.PostView, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var matches []database.PostSearch
	var result []models.PostView

	matches, err = database.FilterPosts(filters+" GROUP BY p.id;", ctx, tx, args...)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	for _, match := range matches {
		var post models.PostView
		post.ID = match.ID
		post.Poster = postUser(match.PosterID, match.Poster, match.PosterAvatar)
		post.Title = match.Title
		post.Body = match.Body
		post.Image = match.Image
		post.Created = match.Created
		post.PostLikers = postLikeSplitter(match.PostLikers)
		post.PostDislikers = postLikeSplitter(match.PostDislikers)
		post.Comments = postComments(match.Comments, match.CommentLikers, match.CommentDislikers)
		post.Categories = postCats(match.PostCategories)
		result = append(result, post)
	}

	return result, err
}

func postUser(id int, userName string, userAvatar string) models.UserView {
	var user models.UserView
	user.ID = id
	user.Username = userName
	user.ProfilePicture = userAvatar
	return user
}

func postComments(comments string, commentLikers string, commentDislikers string) []models.CommentView {
	if comments == "" {
		return nil
	}
	commentGroups := strings.Split(comments, "\x1e")
	var result []models.CommentView
	for _, group := range commentGroups {
		var comment models.CommentView
		parts := strings.SplitN(group, "\x1f", 5)
		commentId, _ := strconv.Atoi(parts[0])
		comment.ID = commentId
		commenterId, _ := strconv.Atoi(parts[1])
		comment.Commenter = postUser(commenterId, parts[2], parts[3])
		comment.Body = parts[4]
		comment.CommentLikers = commentLikeSplitter(commentLikers, commentId)
		comment.CommentDislikers = commentLikeSplitter(commentDislikers, commentId)
		result = append(result, comment)
	}
	return result
}

func postLikeSplitter(parts string) []models.UserView {
	if parts == "" {
		return nil
	}
	likersGroups := strings.Split(parts, "\x1e")
	var allLikes []models.UserView
	for _, likers := range likersGroups {
		splits := strings.SplitN(likers, "\x1f", 3)
		id, _ := strconv.Atoi(splits[0])
		allLikes = append(allLikes, postUser(id, splits[1], splits[2]))
	}
	return allLikes
}

func commentLikeSplitter(parts string, matchID int) []models.UserView {
	if parts == "" {
		return nil
	}
	likersGroups := strings.Split(parts, "\x1e")
	var allLikes []models.UserView
	for _, likers := range likersGroups {
		splits := strings.SplitN(likers, "\x1f", 4)
		id, _ := strconv.Atoi(splits[0])
		if id != matchID {
			continue
		}
		likersId, _ := strconv.Atoi(splits[1])
		allLikes = append(allLikes, postUser(likersId, splits[2], splits[3]))
	}
	return allLikes
}

func postCats(match string) []models.CategoryView {
	if match == "" {
		return nil
	}
	var categories []models.CategoryView
	splittedCat := strings.Split(match, "\x1e")
	for _, catString := range splittedCat {
		var category models.CategoryView
		catAndType := strings.SplitN(catString, "\x1f", 3)
		id, _ := strconv.Atoi(catAndType[0])
		category.ID = id
		category.Category = catAndType[1]
		category.Type = catAndType[2]
		categories = append(categories, category)
	}
	return categories
}

func CreatePost(cookie *http.Cookie, newPost models.Post, db *sql.DB) (error, int) {

	if newPost.Body == "" || newPost.Title == "" {
		return errors.New("Empty post"), http.StatusBadRequest // do nothing
	}

	tx, err := db.Begin()
	if err != nil {
		return err, http.StatusInternalServerError
	}
	defer tx.Rollback()

	userID, err := database.GetUserIDWithToken(tx, cookie.Value)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	newPost.UserID, newPost.Created = userID, time.Now()

	err = database.CreatePost(tx, newPost)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	err = tx.Commit()
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusSeeOther
}
