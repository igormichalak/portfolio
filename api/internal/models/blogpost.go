package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BlogPost struct {
	ID         int       `json:"id"`
	Slug       string    `json:"slug"`
	Title      string    `json:"title"`
	Body       []byte    `json:"-"`
	ParsedBody string    `json:"body"`
	Tags       []BlogTag `json:"tags,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type BlogPostFeedEntry struct {
	ID        int       `json:"id"`
	Slug      string    `json:"slug"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BlogPostModel struct {
	DB *pgxpool.Pool
}

func (m *BlogPostModel) Get(id int) (BlogPost, error) {
	stmt := `SELECT id, slug, title, body, created, updated FROM blog_posts WHERE id = $1`

	p := BlogPost{}

	row := m.DB.QueryRow(context.Background(), stmt, id)
	err := row.Scan(&p.ID, &p.Slug, &p.Title, &p.Body, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return BlogPost{}, err
	}

	p.ParsedBody = `<p>Article content...</p>`

	return p, nil
}

func (m *BlogPostModel) GetFeedPosts() ([]BlogPostFeedEntry, error) {
	stmt := "SELECT id, slug, title, created, updated FROM blog_posts ORDER BY updated"

	rows, err := m.DB.Query(context.Background(), stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var postEntries []BlogPostFeedEntry

	for rows.Next() {
		p := BlogPostFeedEntry{}

		err := rows.Scan(&p.ID, &p.Slug, &p.Title, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}

		postEntries = append(postEntries, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return postEntries, nil
}
