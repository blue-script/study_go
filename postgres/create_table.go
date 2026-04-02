package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, con *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS books(
		id SERIAL PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		author VARCHAR(60) NOT NULL,
		review TEXT,
		release_year INTEGER NOT NULL,
		is_read BOOLEAN NOT NULL DEFAULT FALSE,
		added_at TIMESTAMP NOT NULL,
		completed_at TIMESTAMP 
	)
	`

	_, err := con.Exec(ctx, sqlQuery)

	return err
}
