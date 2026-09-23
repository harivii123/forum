package service

import (
	"strings"
)

func Search(q string) (param string, arg string) {
	return " WHERE post_fts MATCH ?", `"` + strings.ReplaceAll(q, `"`, `""`) + `"`
}
