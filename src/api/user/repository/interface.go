package repository

import (
	"context"

	"example.com/sample-go-wire/src/api/models"
)

type UserRepository interface {
	Add(context.Context, models.User) error
	Find(ctx context.Context, name string) (*models.User, error)
}
