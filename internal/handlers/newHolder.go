package handlers

import (
	"database/sql"
	"forum/internal/template"
)
	

type Holder struct { //can be used to carry other stuff too if needed
	db *sql.DB
	engine *template.Engine
}

func NewHolder(db *sql.DB, engine *template.Engine) *Holder{
	return &Holder{db: db, engine: engine}
}
