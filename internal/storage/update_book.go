package storage

import (
	"context"
	"fmt"
)

type UpdateBookQuery struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Authors     []int  `json:"authors"`
}

func (s *Storage) UpdateBook(ctx context.Context, b *UpdateBookQuery) error {
	_, err := s.db.Query(ctx, "UPDATE books SET name = $2, description = $3, authors = $4 WHERE id = $1", b.ID)

	if err != nil {
		return fmt.Errorf("error scan, %w", err)
	}

	//var rows [][]interface{}
	//for _, authorID := range b.Authors {
	//	rows = append(rows, []interface{}{
	//		bookID,
	//		authorID,
	//	})
	//}
	//
	//_, copyFromErr := s.db.CopyFrom(
	//	ctx,
	//	pgx.Identifier{"book_author"},
	//	[]string{"book_id", "author_id"},
	//	pgx.CopyFromRows(rows),
	//)
	//if copyFromErr != nil {
	//	return 0, fmt.Errorf("query error, %w", copyFromErr)
	//}

	return nil
}
