package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BlogPost struct {
	ID      int
	Slug    string
	Title   string
	Body    []byte
	Tags    []BlogTag
	Created time.Time
	Updated time.Time
}

type BlogPostModel struct {
	DBPool *pgxpool.Pool
}
