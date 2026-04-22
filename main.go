package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/blue-script/study_go/handlers"
	"github.com/blue-script/study_go/postgres"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	con, err := postgres.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}
	defer func(con *pgx.Conn, ctx context.Context) {
		err := con.Close(ctx)
		if err != nil {
			fmt.Println(err)
		}
	}(con, ctx)

	if err := postgres.CreateTable(ctx, con); err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	h := handlers.NewEmployeeHandler(con)
	mux.HandleFunc("GET /employees", h.Handle)
	mux.HandleFunc("POST /employees", h.Handle)
	mux.HandleFunc("DELETE /employees", h.Handle)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
