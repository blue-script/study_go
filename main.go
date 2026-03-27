package main

import (
	"context"
	"fmt"

	"github.com/blue-script/study_go/postgres"
)

func main() {
	ctx := context.Background()

	con, err := postgres.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := con.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Ok")
}
