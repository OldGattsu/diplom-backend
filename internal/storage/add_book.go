package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type AddBookQuery struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Poster      string `json:"poster"`
	Authors     []int  `json:"authors"`
}

func (s *Storage) AddBook(ctx context.Context, b *AddBookQuery) (int, error) {
	row := s.db.QueryRow(ctx, "INSERT INTO books (name, description, poster) VALUES ($1, $2, $3) RETURNING id;", b.Name, b.Description, b.Poster)

	var bookID int
	errScan := row.Scan(&bookID)
	if errScan != nil {
		return 0, fmt.Errorf("error scan, %w", errScan)
	}

	var rows [][]interface{}
	for _, authorID := range b.Authors {
		rows = append(rows, []interface{}{
			bookID,
			authorID,
		})
	}

	_, copyFromErr := s.db.CopyFrom(
		ctx,
		pgx.Identifier{"book_author"},
		[]string{"book_id", "author_id"},
		pgx.CopyFromRows(rows),
	)
	if copyFromErr != nil {
		return 0, fmt.Errorf("query error, %w", copyFromErr)
	}

	return bookID, nil
}
