package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func SelectRows(ctx context.Context, con *pgx.Conn) ([]EmployeeModel, error) {
	sqlQuery := `
		SELECT id, full_name, position
		FROM employees
	`

	rows, err := con.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employees := make([]EmployeeModel, 0)

	for rows.Next() {
		var employee EmployeeModel

		err := rows.Scan(
			&employee.ID,
			&employee.FullName,
			&employee.Position,
		)

		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return employees, nil
}
