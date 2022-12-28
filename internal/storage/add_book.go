package storage

import (
	"context"
	"fmt"

	"github.com/oldgattsu/diplom2/internal/models"
)

func (s *Storage) AddBook(ctx context.Context, b *models.Book) (int, error) {
	// todo: implement id sequence
	_, err := s.db.Exec(ctx, "INSERT INTO books (id, name) VALUES ($1, $2)", 10, b.Name)
	if err != nil {
		return 0, fmt.Errorf("query error, %w", err)
	}

	return 10, nil
}
