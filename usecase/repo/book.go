package repo

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ntheanh201/aws-lambda-go-mongodb/models"
)

type BookRepo struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewBookRepo(client *mongo.Client) *BookRepo {
	return &BookRepo{
		client: client,
		coll:   client.Database(os.Getenv("DB_NAME")).Collection("books"),
	}
}

func (b *BookRepo) GetAllBooks() ([]models.Book, error) {
	cursor, err := b.coll.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal("Repo: ", err)
	}
	var books []models.Book
	if err = cursor.All(context.TODO(), &books); err != nil {
		log.Fatal("Parsing: ", err)
	}
	return books, nil
}
