package main

import (
	"context"
	"fmt"
	"time"

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

	if err := postgres.CreateTable(ctx, con); err != nil {
		panic(err)
	}

	review := "Five"
	timeNow := time.Time{}

	if err := postgres.InsertRow(
		ctx,
		con,
		postgres.BookModel{
			Title:       "Five",
			Author:      "Five",
			Review:      &review,
			ReleaseYear: 1994,
			IsRead:      true,
			AddedAt:     timeNow,
			CompletedAt: &timeNow,
		}); err != nil {
		panic(err)
	}

	// for _, book := range books {
	// 	if book.Id == 3 {
	// 		review := "where"
	// 		now := time.Now()
	// 		book.Title = "where"
	// 		book.Author = "where"
	// 		book.Review = &review
	// 		book.ReleaseYear = 2026
	// 		book.IsRead = true
	// 		book.AddedAt = now
	// 		book.CompletedAt = &now

	// 		if err := postgres.UpdateTask(ctx, con, book); err != nil {
	// 			panic(err)
	// 		}
	// 	}
	// }

	// ids := []int{1, 2, 3, 4}
	// if err := postgres.DeleteBooks(ctx, con, ids); err != nil {
	// 	panic(err)
	// }

	err = postgres.ListPage(ctx, con, 5)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}
