package storage

import "context"

func (s *Storage) SaveToken(ctx context.Context, userID int, token string) error {
	_, err := s.db.Exec(ctx, "INSERT INTO tokens (user_id, token) VALUES($1, $2)", userID, token)
	return err
}
