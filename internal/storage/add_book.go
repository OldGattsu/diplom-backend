package storage

import (
	"context"
	"fmt"
	"github.com/oldgattsu/diplom2/internal/models"
)

func (s *Storage) AddBook(ctx context.Context, b *models.Book) (int, error) {
	row := s.db.QueryRow(ctx, "INSERT INTO books (name) VALUES ($1) RETURNING id;", b.Name)

	var id int
	errScan := row.Scan(&id)
	if errScan != nil {
		return 0, fmt.Errorf("error scan, %w", errScan)
	}

	return id, nil
}
