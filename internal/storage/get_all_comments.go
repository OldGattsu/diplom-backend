package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"

	"github.com/oldgattsu/diplom2/internal/models"
)

func (s *Storage) GetAllComments(ctx context.Context) ([]*models.Comment, error) {
	rows, errQuery := s.db.Query(ctx, "SELECT id, user_id, book_id, text FROM comments")
	if errQuery != nil {
		return nil, fmt.Errorf("query error, %w", errQuery)
	}

	var res []*models.Comment

	for rows.Next() {
		c := &models.Comment{}
		errScan := rows.Scan(&c.ID, &c.UserId, &c.BookID, &c.Text)
		if errScan != nil {
			return nil, fmt.Errorf("scan error, %w", errScan)
		}

		userRow := s.db.QueryRow(ctx, "SELECT id, name FROM users where id = $1", c.UserId)
		u := &models.User{}
		errUserScan := userRow.Scan(&u.ID, &u.Name)
		if errUserScan != nil {
			if errors.Is(errScan, pgx.ErrNoRows) {
				return nil, ErrCommentNotFound
			}
			return nil, fmt.Errorf("scan error, %w", errScan)
		}
		c.User = u

		bookRow := s.db.QueryRow(ctx, "SELECT id, name FROM books where id = $1", c.BookID)
		b := &models.Book{}
		errBookScan := bookRow.Scan(&b.ID, &b.Name)
		if errBookScan != nil {
			if errors.Is(errScan, pgx.ErrNoRows) {
				return nil, ErrCommentNotFound
			}
			return nil, fmt.Errorf("scan error, %w", errScan)
		}
		c.Book = b

		res = append(res, c)
	}

	return res, nil
}
