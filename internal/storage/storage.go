package storage

import (
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
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
