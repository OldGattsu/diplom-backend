package storage

import (
	"context"
	"fmt"

	"github.com/oldgattsu/diplom2/internal/models"
)

func (s *Storage) GetUsers(ctx context.Context) ([]*models.User, error) {
	rows, errQuery := s.db.Query(ctx, "SELECT id, name, is_blocked FROM users")
	if errQuery != nil {
		return nil, fmt.Errorf("query error, %w", errQuery)
	}

	var res []*models.User

	for rows.Next() {
		u := &models.User{}

		errScan := rows.Scan(&u.ID, &u.Name, &u.IsBlocked)
		if errScan != nil {
			return nil, fmt.Errorf("scan error, %w", errScan)
		}

		res = append(res, u)
	}

	return res, nil
}
