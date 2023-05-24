package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/ntheanh201/aws-lambda-go-mongodb/db"
	"github.com/ntheanh201/aws-lambda-go-mongodb/usecase"
	"github.com/ntheanh201/aws-lambda-go-mongodb/usecase/repo"
)

func handler(ctx context.Context, event events.CloudWatchEvent) error {
	client, err := db.ConnectMongoDB()

	if err != nil {
		panic(err)
	}

	fmt.Println("cloudWatchEvent: ", event)
	fmt.Println("Lambda function executed")

	bookUseCase := usecase.NewBookUseCase(repo.NewBookRepo(client))
	books, _ := bookUseCase.GetAllBooks(ctx)
	if err != nil {
		panic(err)
	}
	for _, book := range books {
		println("Name: ", book.Name, " - Author: ", book.Author)
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
