package service

import (
	"context"

	domain "example.com/sample-go-wire/src/domain/client"
	"example.com/sample-go-wire/src/usecase"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

type client struct {
	repo domain.ClientRepository
}

func New(repo domain.ClientRepository) domain.ClientService {
	wire.Bind(new(domain.ClientService), new(*client))
	return &client{
		repo: repo,
	}
}

func (c *client) Add(ctx context.Context, client domain.Client) error {
	if len(client.Name) > 20 {
		return usecase.ErrInvalidNameLength
	}

	clientExists, err := c.repo.Find(ctx, client.Name)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return err
		}
	}

	if clientExists != nil {
		return usecase.ErrClientAlreadyExists
	}

	return c.repo.Add(ctx, client)
}
