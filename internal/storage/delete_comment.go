package storage

import (
	"context"
	"fmt"
)

func (s *Storage) DeleteComment(ctx context.Context, commentID int) error {
	_, deleteCommentErr := s.db.Query(ctx, "DELETE FROM comments WHERE id = $1", commentID)
	if deleteCommentErr != nil {
		return fmt.Errorf("delete comment error, %w", deleteCommentErr)
	}

	return nil
}
