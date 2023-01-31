package storage

import (
	"context"
	"fmt"

	"github.com/oldgattsu/diplom2/internal/models"
)

func (s *Storage) GetComments(ctx context.Context, bookID int) ([]*models.Comment, error) {
	rows, errQuery := s.db.Query(ctx, "SELECT id, user_id, text, book_id FROM comments WHERE book_id = $1", bookID)
	if errQuery != nil {
		return nil, fmt.Errorf("query error, %w", errQuery)
	}

	var res []*models.Comment

	for rows.Next() {
		b := &models.Comment{}

		errScan := rows.Scan(&b.ID, &b.UserId, &b.Text, &b.BookID)
		if errScan != nil {
			return nil, fmt.Errorf("scan error, %w", errScan)
		}

		res = append(res, b)
	}

	return res, nil
}
