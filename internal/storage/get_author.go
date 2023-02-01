package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"

	"github.com/oldgattsu/diplom2/internal/models"
)

func (s *Storage) GetAuthor(ctx context.Context, authorID int) (*models.Author, error) {
	row := s.db.QueryRow(ctx, "SELECT id, name, description FROM authors WHERE id = $1", authorID)

	b := &models.Author{}

	errScan := row.Scan(&b.ID, &b.Name, &b.Description)
	if errScan != nil {
		if errors.Is(errScan, pgx.ErrNoRows) {
			return nil, ErrAuthorNotFound
		}
		return nil, fmt.Errorf("scan error, %w", errScan)
	}

	books, errGetBooks := s.db.Query(ctx, "SELECT a.id, a.name FROM book_author AS ba LEFT JOIN books AS a on a.id = ba.book_id WHERE ba.author_id = $1", authorID)
	if errGetBooks != nil {
		return nil, fmt.Errorf("error get books, %w", errGetBooks)
	}

	for books.Next() {
		a := &models.Book{}

		errScanBook := books.Scan(&a.ID, &a.Name)
		if errGetBooks != nil {
			return nil, fmt.Errorf("error scan book, %w", errScanBook)
		}

		b.Books = append(b.Books, a)
	}

	return b, nil
}
