package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, con *pgx.Conn, employee EmployeeModel) error {
	sqlQuery := `
		INSERT INTO employees(full_name, position)
		VALUES($1, $2)
	`

	fmt.Println(employee)
	fmt.Println(employee.FullName)
	fmt.Println(employee.Position)
	_, err := con.Exec(
		ctx,
		sqlQuery,
		employee.FullName,
		employee.Position,
	)

	return err
}
