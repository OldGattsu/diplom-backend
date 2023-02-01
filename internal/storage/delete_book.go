package storage

import (
	"context"
	"fmt"
)

func (s *Storage) DeleteBook(ctx context.Context, bookID int) error {
	_, deleteBookErr := s.db.Query(ctx, "DELETE FROM books WHERE id = $1", bookID)
	if deleteBookErr != nil {
		return fmt.Errorf("delete book error, %w", deleteBookErr)
	}

	_, deleteBookCommentsErr := s.db.Query(ctx, "DELETE FROM comments WHERE book_id = $1", bookID)
	if deleteBookCommentsErr != nil {
		return fmt.Errorf("delete comments error, %w", deleteBookCommentsErr)
	}

	return nil
}
