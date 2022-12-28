package storage

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"

	"github.com/oldgattsu/diplom2/internal/models"
)

var ErrUserNotFound = errors.New("user not found")

type Storage struct {
	db     *pgxpool.Pool
	logger *zap.Logger
}

func New(db *pgxpool.Pool, logger *zap.Logger) *Storage {
	s := &Storage{
		db:     db,
		logger: logger,
	}

	return s
}

func (s *Storage) SaveToken(ctx context.Context, userID int, token string) error {
	_, err := s.db.Exec(ctx, "INSERT INTO tokens (user_id, token) VALUES($1, $2)", userID, token)
	return err
}

func (s *Storage) GetUser(ctx context.Context, email, password string) (*models.User, error) {
	row := s.db.QueryRow(ctx, "SELECT id, name, email, password FROM users WHERE email = $1 LIMIT 1", email)

	u := &models.User{}
	var passwordDB string

	errScan := row.Scan(&u.ID, &u.Name, &u.Email, &passwordDB)
	if errScan != nil {
		if errors.Is(errScan, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("error scan, %w", errScan)
	}

	errPassword := bcrypt.CompareHashAndPassword([]byte(passwordDB), []byte(password))
	if errPassword != nil {
		return nil, ErrUserNotFound
	}

	return u, nil
}
