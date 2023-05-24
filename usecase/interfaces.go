package usecase

import (
	"context"

	"github.com/ntheanh201/aws-lambda-go-mongodb/models"
)

type (
	Book interface {
		GetAllBooks(ctx context.Context) ([]models.Book, error)
	}

	BookRepo interface {
		GetAllBooks() ([]models.Book, error)
	}
)
