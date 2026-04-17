package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, con *pgx.Conn, user UserModel) error {
	sqlQuery := `
		INSERT INTO users(full_name, phone_number)
		VALUES($1, $2)
	`
	_, err := con.Exec(
		ctx,
		sqlQuery,
		user.FullName,
		user.PhoneNumber,
	)

	return err
}
