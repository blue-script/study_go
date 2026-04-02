package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateTask(ctx context.Context, con *pgx.Conn, book BookModel) error {
	sqlQuery := `
		UPDATE books
		SET title=$1, author=$2, review=$3, release_year=$4, is_read=$5, added_at=$6, completed_at=$7
		WHERE id=$8
	`

	_, err := con.Exec(ctx, sqlQuery, book.Title, book.Author, book.Review, book.ReleaseYear, book.IsRead, book.AddedAt, book.CompletedAt, book.Id)
	return err
}
