package service

import (
	"context"

	"example.com/sample-go-wire/src/api/models"
)

type ClientService interface {
	Add(context.Context, models.Client) error
}
