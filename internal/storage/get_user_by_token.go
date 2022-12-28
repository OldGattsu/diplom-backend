package storage

import (
	"context"
	_ "embed"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"

	"github.com/oldgattsu/diplom2/internal/models"
)

//go:embed get_user_by_token.sql
var queryGetUserByToken string

func (s *Storage) GetUserByToken(ctx context.Context, token string) (*models.User, error) {
	row := s.db.QueryRow(ctx, queryGetUserByToken, token)

	u := &models.User{}

	errScan := row.Scan(&u.ID, &u.Name, &u.Email)
	if errScan != nil {
		if errors.Is(errScan, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("error scan, %w", errScan)
	}

	return u, nil
}
