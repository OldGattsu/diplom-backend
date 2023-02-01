package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"

	"github.com/oldgattsu/diplom2/internal/models"
)

func (s *Storage) GetComments(ctx context.Context, bookID int) ([]*models.Comment, error) {
	rows, errQuery := s.db.Query(ctx, "SELECT id, user_id, text, book_id FROM comments WHERE book_id = $1", bookID)
	if errQuery != nil {
		return nil, fmt.Errorf("query error, %w", errQuery)
	}

	var res []*models.Comment

	for rows.Next() {
		b := &models.Comment{}

		errScan := rows.Scan(&b.ID, &b.UserId, &b.Text, &b.BookID)
		if errScan != nil {
			return nil, fmt.Errorf("scan error, %w", errScan)
		}

		userRow := s.db.QueryRow(ctx, "SELECT id, name, email FROM users where id = $1", b.UserId)

		u := &models.User{}
		errUserScan := userRow.Scan(&u.ID, &u.Name, &u.Email)
		if errUserScan != nil {
			if errors.Is(errScan, pgx.ErrNoRows) {
				return nil, ErrCommentNotFound
			}
			return nil, fmt.Errorf("scan error, %w", errScan)
		}

		b.User = u

		res = append(res, b)
	}

	return res, nil
}
