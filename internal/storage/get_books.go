package storage

import (
	"context"
	"fmt"

	"github.com/oldgattsu/diplom2/internal/models"
)

func (s *Storage) GetBooks(ctx context.Context) ([]*models.Book, error) {
	rows, errQuery := s.db.Query(ctx, "SELECT id, name FROM books")
	if errQuery != nil {
		return nil, fmt.Errorf("query error, %w", errQuery)
	}

	var res []*models.Book

	for rows.Next() {
		b := &models.Book{}

		errScan := rows.Scan(&b.ID, &b.Name)
		if errScan != nil {
			return nil, fmt.Errorf("scan error, %w", errScan)
		}

		res = append(res, b)
	}

	return res, nil
}
