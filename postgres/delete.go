package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteBooks(ctx context.Context, con *pgx.Conn, booksIds []int) error {
	sqlQuery := `
		DELETE FROM books
		WHERE id = ANY($1)
	`

	_, err := con.Exec(ctx, sqlQuery, booksIds)

	return err
}
