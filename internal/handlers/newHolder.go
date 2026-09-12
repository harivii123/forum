package handlers

import "database/sql"

type Holder struct { //can be used to carry other stuff too if needed
	db *sql.DB
}

func NewHolder(db *sql.DB) *Holder{
	return &Holder{db: db}
}
