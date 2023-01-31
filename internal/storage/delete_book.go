package storage

import (
	"context"
	"fmt"
)

func (s *Storage) DeleteBook(ctx context.Context, bookID int) error {
	_, err := s.db.Query(ctx, "DELETE FROM books WHERE id = $1", bookID)

	if err != nil {
		return fmt.Errorf("scan error, %w", err)
	}

	return nil
}
