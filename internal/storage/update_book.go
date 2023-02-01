package storage

import (
	"context"
	"fmt"
)

type UpdateBookQuery struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Authors     []int  `json:"authors"`
}

func (s *Storage) UpdateBook(ctx context.Context, b *UpdateBookQuery) error {
	_, err := s.db.Exec(ctx, "UPDATE books SET name = $2, description = $3 WHERE id = $1", b.ID, b.Name, b.Description)

	if err != nil {
		return fmt.Errorf("error update book, %w", err)
	}

	return nil
}
