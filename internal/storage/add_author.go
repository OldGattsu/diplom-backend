package storage

import (
	"context"
	"fmt"

	"github.com/oldgattsu/diplom2/internal/models"
)

func (s *Storage) AddAuthor(ctx context.Context, a *models.Author) (int, error) {
	row := s.db.QueryRow(ctx, "INSERT INTO authors (name, description) VALUES ($1, $2) RETURNING id;", a.Name, a.Description)

	var id int
	errScan := row.Scan(&id)
	if errScan != nil {
		return 0, fmt.Errorf("error scan, %w", errScan)
	}

	return id, nil
}
