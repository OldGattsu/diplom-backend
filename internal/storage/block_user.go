package storage

import (
	"context"
	"fmt"
)

func (s *Storage) BlockUser(ctx context.Context, id int, isBlocked bool) error {
	_, err := s.db.Exec(ctx, "UPDATE users SET is_blocked = $2 WHERE id = $1", id, isBlocked)

	if err != nil {
		return fmt.Errorf("error block user, %w", err)
	}

	return nil
}
