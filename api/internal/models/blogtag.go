package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BlogTag struct {
	ID   int
	Name string
}

type BlogTagModel struct {
	DB *pgxpool.Pool
}

func (m *BlogTagModel) GetAll() ([]BlogTag, error) {
	stmt := "SELECT (id, name) FROM blog_tags"

	rows, err := m.DB.Query(context.Background(), stmt)
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
