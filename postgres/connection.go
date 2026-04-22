package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	db := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, db)

	return pgx.Connect(ctx, connString)
}
