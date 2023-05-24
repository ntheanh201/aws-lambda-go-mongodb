package usecase

import (
	"context"

	"github.com/ntheanh201/aws-lambda-go-mongodb/models"
)

type BookUseCase struct {
	r BookRepo
}

func NewBookUseCase(repo BookRepo) *BookUseCase {
	return &BookUseCase{r: repo}
}

func (b *BookUseCase) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	books, _ := b.r.GetAllBooks()
	return books, nil
}
