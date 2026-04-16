package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	postgres_connect := os.Getenv("postgres_connect")
	if postgres_connect == "" {
		fmt.Println("Can't connect")
	} else {
		fmt.Println("Success connect")
	}
	return pgx.Connect(ctx, postgres_connect)
}
