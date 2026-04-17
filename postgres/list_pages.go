package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func ListPage(ctx context.Context, con *pgx.Conn) error {
	sqlQuery := `
		SELECT full_name, phone_number FROM users
		ORDER BY id
	`
	row, err := con.Query(ctx, sqlQuery)
	if err != nil {
		return err
	}

	var users []UserModel

	for row.Next() {
		var user UserModel

		err := row.Scan(&user.FullName, &user.PhoneNumber)
		if err != nil {
			row.Close()
			return err
		}

		users = append(users, user)
	}

	rowErr := row.Err()

	row.Close()

	if rowErr != nil {
		return rowErr
	}

	fmt.Printf("%v\n", users)

	return nil
}
