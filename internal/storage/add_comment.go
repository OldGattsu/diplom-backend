package storage

import (
	"context"
	"fmt"
)

type AddCommentQuery struct {
	UserID int    `json:"user_id"`
	BookID int    `json:"book_id"`
	Text   string `json:"text"`
}

func (s *Storage) AddComment(ctx context.Context, c *AddCommentQuery) (int, error) {
	row := s.db.QueryRow(ctx, "INSERT INTO comments (user_id, book_id, text) VALUES ($1, $2, $3) RETURNING id;", c.UserID, c.BookID, c.Text)

	var id int
	errScan := row.Scan(&id)
	if errScan != nil {
		return 0, fmt.Errorf("error scan, %w", errScan)
	}

	return id, nil
}
