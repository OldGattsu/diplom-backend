package storage

import (
	"context"
	"fmt"

	"github.com/oldgattsu/diplom2/internal/models"
)

func (s *Storage) GetAuthors(ctx context.Context) ([]*models.Author, error) {
	rows, errQuery := s.db.Query(ctx, "SELECT id, name FROM authors")
	if errQuery != nil {
		return nil, fmt.Errorf("query error, %w", errQuery)
	}

	var res []*models.Author

	for rows.Next() {
		b := &models.Author{}

		errScan := rows.Scan(&b.ID, &b.Name)
		if errScan != nil {
			return nil, fmt.Errorf("scan error, %w", errScan)
		}

		res = append(res, b)
	}

	return res, nil
}
