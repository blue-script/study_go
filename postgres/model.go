package postgres

type EmployeeModel struct {
	ID       int     `db:"id"`
	FullName string  `db:"full_name"`
	Position *string `db:"position"`
}
