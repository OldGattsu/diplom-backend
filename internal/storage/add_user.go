package storage

import (
	"context"
	"fmt"
	"github.com/oldgattsu/diplom2/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type AddUserQuery struct {
	Name     string
	Email    string
	Password string
	IsAdmin  bool
}

func (s *Storage) AddUser(ctx context.Context, uq *AddUserQuery) (*models.User, error) {
	hash, errHash := bcrypt.GenerateFromPassword([]byte(uq.Password), bcrypt.DefaultCost)
	if errHash != nil {
		return nil, fmt.Errorf("error hash generate, %w", errHash)
	}

	row := s.db.QueryRow(ctx, "INSERT INTO users (name, email, password, is_admin) VALUES ($1, $2, $3, $4) RETURNING id, name, email, is_admin;", uq.Name, uq.Email, hash, uq.IsAdmin)

	u := &models.User{}
	errScan := row.Scan(&u.ID, &u.Name, &u.Email, &u.IsAdmin)
	fmt.Printf("user: %v", u)
	// todo: check error обработать ошибку нормально
	if errScan != nil {
		return nil, fmt.Errorf("error scan, %w", errScan)
	}

	return u, nil
}
