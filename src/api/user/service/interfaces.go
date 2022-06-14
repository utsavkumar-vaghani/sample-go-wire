package service

import (
	"context"

	"example.com/sample-go-wire/src/api/models"
)

type UserService interface {
	Add(context.Context, models.User) error
}
