package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func SelectRows(ctx context.Context, con *pgx.Conn) ([]BookModel, error) {
	sqlQuery := `
		SELECT id, title, author, review, release_year, is_read, added_at, completed_at
		FROM books
		LIMIT 5
		OFFSET 11
	`

	rows, err := con.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]BookModel, 0)

	for rows.Next() {
		var book BookModel

		err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.Author,
			&book.Review,
			&book.ReleaseYear,
			&book.IsRead,
			&book.AddedAt,
			&book.CompletedAt,
		)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}
