package models

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BlogPost struct {
	ID            int       `json:"id"`
	Slug          string    `json:"slug"`
	Title         string    `json:"title"`
	Body          []byte    `json:"-"`
	ParsedBody    string    `json:"body"`
	Tags          []BlogTag `json:"tags,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	IsCodeSnippet bool      `json:"is_code_snippet"`
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
	stmt := `SELECT id, slug, title, body, created, updated, is_code_snippet 
	FROM blog_posts 
	WHERE id = $1`

	p := BlogPost{}

	row := m.DB.QueryRow(context.Background(), stmt, id)
	err := row.Scan(&p.ID, &p.Slug, &p.Title, &p.Body, &p.CreatedAt, &p.UpdatedAt, &p.IsCodeSnippet)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return BlogPost{}, ErrNoRecord
		}
		return BlogPost{}, err
	}

	tags, err := m.GetPostTags(id)
	if err != nil {
		tags = []BlogTag{}
	}

	p.Tags = tags
	p.ParsedBody = `<p>Article content...</p>`

	return p, nil
}

func (m *BlogPostModel) GetBySlug(slug string) (BlogPost, error) {
	stmt := `SELECT id, slug, title, body, created, updated, is_code_snippet 
	FROM blog_posts 
	WHERE slug = $1`

	p := BlogPost{}

	row := m.DB.QueryRow(context.Background(), stmt, slug)
	err := row.Scan(&p.ID, &p.Slug, &p.Title, &p.Body, &p.CreatedAt, &p.UpdatedAt, &p.IsCodeSnippet)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return BlogPost{}, ErrNoRecord
		}
		return BlogPost{}, err
	}

	tags, err := m.GetPostTags(p.ID)
	if err != nil {
		tags = []BlogTag{}
	}

	p.Tags = tags
	p.ParsedBody = `<p>Article content...</p>`

	return p, nil
}

func (m *BlogPostModel) GetPostTags(id int) ([]BlogTag, error) {
	stmt := `SELECT id, name 
	FROM blog_tags 
	WHERE id IN (SELECT tag_id FROM post_tags WHERE post_id = $1)`

	rows, err := m.DB.Query(context.Background(), stmt, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []BlogTag

	for rows.Next() {
		t := BlogTag{}

		if err := rows.Scan(&t.ID, &t.Name); err != nil {
			return nil, err
		}

		tags = append(tags, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}

func (m *BlogPostModel) GetFeedPosts(onlyCodeSnippets bool) ([]BlogPostFeedEntry, error) {
	stmt := `SELECT id, slug, title, created, updated 
	FROM blog_posts 
	WHERE is_code_snippet = $1
	ORDER BY updated DESC`

	rows, err := m.DB.Query(context.Background(), stmt, onlyCodeSnippets)
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
