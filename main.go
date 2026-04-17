package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

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

	NEW_USER := os.Getenv("NEW_USER")
	if NEW_USER == "YES" {
		reader := bufio.NewReader(os.Stdin)

		userName, _ := reader.ReadString('\n')

		var phoneInput string
		var phoneNumber *string
		phoneInput, _ = reader.ReadString('\n')
		fmt.Println(phoneInput)
		if phoneInput == "\n" {
			phoneNumber = nil
		} else {
			phoneNumber = &phoneInput
		}

		user := postgres.UserModel{
			FullName:    userName,
			PhoneNumber: phoneNumber,
		}
		if err := postgres.InsertRow(ctx, con, user); err != nil {
			fmt.Println(err)
		}
		fmt.Println("User successful add")
	} else if NEW_USER == "NO" {
		if err := postgres.ListPage(ctx, con); err != nil {
			fmt.Println("Error read table user")
		}
	} else {
		fmt.Println("Incorrect NEW_USER")
	}

	fmt.Println("Success")
}
