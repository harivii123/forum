package service

import (
	"database/sql"
	"forum/internal/database"
	"forum/internal/models"
	"strconv"
	"strings"
)

// takes value from main search bar and finds all matches using fts(fast text search)
func MainSearch(mainSearchValue string, db *sql.DB) ([]models.PostView, error) {
	//if "" nothing happens
	if mainSearchValue == "" {
		// log.Println(4)
		return nil, nil
	}

	tx, err := db.Begin()
	if err != nil {
		// log.Println(5)
		return nil, err
	}
	defer tx.Rollback()

	var matches []database.PostSearch
	var result []models.PostView

	matches, err = database.FindAMatch(tx, mainSearchValue)
	if err != nil {
		// log.Println(6)
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		// log.Println(7)
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

	// log.Println(match, "here")
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

func postCats(match string) []models.PostCategory {
	if match == "" {
		return nil
	}
	var categories []models.PostCategory
	splittedCat := strings.Split(match, "\x1f")
	for _, category := range splittedCat {
		cat := models.PostCategory(category)
		categories = append(categories, cat)
	}
	return categories
}
