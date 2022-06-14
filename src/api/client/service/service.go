package service

import (
	"context"

	"example.com/sample-go-wire/src/api/client/repository"
	"example.com/sample-go-wire/src/api/errors"
	"example.com/sample-go-wire/src/api/models"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

type client struct {
	repo repository.ClientRepository
}

func New(repo repository.ClientRepository) ClientService {
	wire.Bind(new(ClientService), new(*client))
	return &client{
		repo: repo,
	}
}

func (c *client) Add(ctx context.Context, client models.Client) error {
	if len(client.Name) > 20 {
		return errors.ErrInvalidNameLength
	}

	clientExists, err := c.repo.Find(ctx, client.Name)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return err
		}
	}

	if clientExists != nil {
		return errors.ErrClientAlreadyExists
	}

	return c.repo.Add(ctx, client)
}
