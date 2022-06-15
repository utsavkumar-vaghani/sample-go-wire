package service

import (
	"context"

	domain "example.com/sample-go-wire/src/domain/user"
	"example.com/sample-go-wire/src/usecase"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

type user struct {
	repo domain.UserRepository
}

func New(repo domain.UserRepository) domain.UserService {
	wire.Bind(new(domain.UserService), new(*user))
	return &user{
		repo: repo,
	}
}

func (u *user) Add(ctx context.Context, user domain.User) error {
	if len(user.Name) > 20 {
		return usecase.ErrInvalidNameLength
	}

	userExists, err := u.repo.Find(ctx, user.Name)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return err
		}
	}

	if userExists != nil {
		return usecase.ErrUserAlreadyExists
	}

	return u.repo.Add(ctx, user)
}
