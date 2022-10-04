package models

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type BlogTag struct {
	ID   int
	Name string
}

type BlogTagModel struct {
	DBPool *pgxpool.Pool
}
