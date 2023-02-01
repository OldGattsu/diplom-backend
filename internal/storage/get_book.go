package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"

	"github.com/oldgattsu/diplom2/internal/models"
)

func (s *Storage) GetBook(ctx context.Context, bookID int) (*models.Book, error) {
	row := s.db.QueryRow(ctx, "SELECT id, name, description, poster FROM books WHERE id = $1", bookID)

	b := &models.Book{}

	errScan := row.Scan(&b.ID, &b.Name, &b.Description, &b.Poster)
	if errScan != nil {
		if errors.Is(errScan, pgx.ErrNoRows) {
			return nil, ErrBookNotFound
		}
		return nil, fmt.Errorf("scan error, %w", errScan)
	}

	authors, errGetAuthors := s.db.Query(ctx, "SELECT a.id, a.name FROM book_author AS ba LEFT JOIN authors AS a on a.id = ba.author_id WHERE ba.book_id = $1", bookID)
	if errGetAuthors != nil {
		return nil, fmt.Errorf("error get authors, %w", errGetAuthors)
	}

	for authors.Next() {
		a := &models.Author{}

		errScanAuthor := authors.Scan(&a.ID, &a.Name)
		if errGetAuthors != nil {
			return nil, fmt.Errorf("error scan author, %w", errScanAuthor)
		}

		b.Authors = append(b.Authors, a)
	}

	return b, nil
}
