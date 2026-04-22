package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteEmployees(ctx context.Context, con *pgx.Conn, employeesIds []int) error {
	sqlQuery := `
		DELETE FROM employees
		WHERE id = ANY($1)
	`

	_, err := con.Exec(ctx, sqlQuery, employeesIds)

	return err
}
