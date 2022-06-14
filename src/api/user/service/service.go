package service

import (
	"context"

	"example.com/sample-go-wire/src/api/errors"
	"example.com/sample-go-wire/src/api/models"
	"example.com/sample-go-wire/src/api/user/repository"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

type user struct {
	repo repository.UserRepository
}

func New(repo repository.UserRepository) UserService {
	wire.Bind(new(UserService), new(*user))
	return &user{
		repo: repo,
	}
}

func (u *user) Add(ctx context.Context, user models.User) error {
	if len(user.Name) > 20 {
		return errors.ErrInvalidNameLength
	}

	userExists, err := u.repo.Find(ctx, user.Name)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return err
		}
	}

	if userExists != nil {
		return errors.ErrUserAlreadyExists
	}

	return u.repo.Add(ctx, user)
}
