package service

import (
	"fmt"
	"strings"
)

func Search(q string) (param string, arg string) {
	return "post_fts MATCH ?", `"` + strings.ReplaceAll(q, `"`, `""`) + `"`
}

func Posts(userID int) (param string, arg int) {
	return "p.user_id = ?", userID
}

func Likes(userID int) (param string, arg int) {
	return "vu.id = ?", userID
}

func Categories(ids []int) (param string, arg []any) {
	for _, id := range ids {
		arg = append(arg, id)
	}
	postCategories := "p.id IN (SELECT post_id FROM category_post WHERE category_id IN "
	catIds := "(" + strings.Repeat("?, ", len(ids)-1) + "?" + ")"
	return fmt.Sprintf(postCategories+catIds+" GROUP BY post_id HAVING COUNT(*) = %d)", len(arg)), arg
}

func Order(order string) string {
	switch order {
	case "Oldest":
		return "ORDER BY p.created ASC"
	case "Liked":
		return "ORDER BY (SELECT COUNT(*) FROM post_vote WHERE post_id = p.id AND vote = 1) DESC"
	default:
		return "ORDER BY p.created DESC"
	}
}
