package database

import (
	"database/sql"
	"log"
	"time"
)

type PostSearch struct {
	ID               int
	Title            string
	PosterID         int
	Body             string
	Image            string
	Created          time.Time
	Poster           string
	PosterAvatar     string
	Comments         string
	CommentLikers    string
	CommentDislikers string
	PostLikers       string
	PostDislikers    string
	PostCategories   string
}

func FindAMatch(tx *sql.Tx) ([]PostSearch, error) {
	// log.Println(searchValue, "rep")
	query := `SELECT p.id, p.title, p.user_id, p.body, COALESCE(p.image, ''), p.created_at, u.username, COALESCE(u.profile_picture, '/static/avatars/default-avatar.png'),
		COALESCE((SELECT group_concat(c.id || char(31) || cu.id || char(31) || cu.username || char(31) || COALESCE(cu.profile_picture, '/static/avatars/default-avatar.png') || char(31) || c.body, char(30))
				FROM comment c JOIN user cu ON cu.id = c.user_id
				WHERE c.post_id = p.id), '') AS comments,
		COALESCE((SELECT group_concat(cv.comment_id || char(31) || cvu.id || char(31) || cvu.username || char(31) || COALESCE(cvu.profile_picture, '/static/avatars/default-avatar.png'), char(30))
				FROM comment_vote cv JOIN user cvu ON cv.user_id = cvu.id JOIN comment c ON c.id = cv.comment_id
				WHERE c.post_id = p.id AND cv.vote = 1), '') AS comment_likers,
		COALESCE((SELECT group_concat(cv.comment_id || char(31) || cvu.id || char(31) || cvu.username || char(31) || COALESCE(cvu.profile_picture, '/static/avatars/default-avatar.png'), char(30))
				FROM comment_vote cv JOIN user cvu ON cv.user_id = cvu.id JOIN comment c ON c.id = cv.comment_id
				WHERE c.post_id = p.id AND cv.vote = -1), '') AS comment_dislikers,
		COALESCE((SELECT group_concat(vu.id || char(31) || vu.username || char(31) || COALESCE(vu.profile_picture, '/static/avatars/default-avatar.png'), char(30))
			FROM post_vote pv JOIN user vu ON vu.id = pv.user_id
			WHERE pv.post_id = p.id AND pv.vote = 1), '') AS likers,
		COALESCE((SELECT group_concat(vu.id || char(31) || vu.username || char(31) || COALESCE(vu.profile_picture, '/static/avatars/default-avatar.png'), char(30))
			FROM post_vote pv JOIN user vu ON vu.id = pv.user_id
			WHERE pv.post_id = p.id AND pv.vote = -1), '') AS dislikers,
		COALESCE((SELECT group_concat(cat.name, char(31))
			FROM category_post cp JOIN category cat ON cat.id = cp.category_id
			WHERE cp.post_id = p.id), '') AS categories
		FROM post p
		JOIN user u ON u.id = p.user_id
		JOIN post_fts fts ON fts.rowid = p.id
		GROUP BY p.id;`
	//WHERE post_fts MATCH ?;
	rows, err := tx.Query(
		// `SELECT u.id, 'user', u.username, 'User Account'
		//  FROM user u
		//  JOIN user_fts fts ON u.id = fts.rowid
		//  WHERE user_fts MATCH ?

		// UNION ALL
		query)
	if err != nil {
		log.Println(1, err)
		return nil, err
	}
	defer rows.Close()

	result := []PostSearch{}

	for rows.Next() {
		var postSearch PostSearch
		err := rows.Scan(&postSearch.ID, &postSearch.Title, &postSearch.PosterID, &postSearch.Body, &postSearch.Image, &postSearch.Created, &postSearch.Poster, &postSearch.PosterAvatar, &postSearch.Comments, &postSearch.CommentLikers, &postSearch.CommentDislikers, &postSearch.PostLikers, &postSearch.PostDislikers, &postSearch.PostCategories)
		if err != nil {
			// log.Println(2)
			return nil, err
		}
		result = append(result, postSearch)
	}
	if rows.Err() != nil {
		// log.Println(3)
		return nil, err
	}

	return result, nil
}
