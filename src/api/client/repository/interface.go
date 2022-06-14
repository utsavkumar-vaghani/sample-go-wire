package repository

import (
	"context"

	"example.com/sample-go-wire/src/api/models"
)

type ClientRepository interface {
	Add(context.Context, models.Client) error
	Find(ctx context.Context, name string) (*models.Client, error)
}
