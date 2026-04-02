package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, con *pgx.Conn, book BookModel) error {
	sqlQuery := `
		INSERT INTO books(title, author, review, release_year, is_read, added_at, completed_at)
		VALUES($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := con.Exec(
		ctx,
		sqlQuery,
		book.Title,
		book.Author,
		book.Review,
		book.ReleaseYear,
		book.IsRead,
		book.AddedAt,
		book.CompletedAt,
	)

	return err
}
