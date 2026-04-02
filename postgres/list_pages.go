package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func ListPage(ctx context.Context, con *pgx.Conn, count int) error {
	sqlQuery := `
		SELECT title, author FROM books
		ORDER BY id
		LIMIT $1 OFFSET $2
	`

	for i := 0; ; i++ {
		offset := i * count
		row, err := con.Query(ctx, sqlQuery, count, offset)
		if err != nil {
			return err
		}

		var books []BookModel

		for row.Next() {
			var book BookModel

			err := row.Scan(&book.Title, &book.Author)
			if err != nil {
				row.Close()
				return err
			}

			books = append(books, book)
		}

		rowErr := row.Err()

		row.Close()

		if rowErr != nil {
			return rowErr
		}

		if len(books) == 0 && i > 0 {
			break
		}

		fmt.Printf("Страница %v: %v\n", i+1, books)

		if len(books) < count {
			break
		}
	}

	return nil
}
