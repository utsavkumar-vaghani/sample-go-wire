package repository

import (
	"context"

	"example.com/sample-go-wire/src/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type client struct {
	db *mongo.Database
}

func New(db *mongo.Database) ClientRepository {
	return &client{
		db: db,
	}
}

func (c client) Add(ctx context.Context, client models.Client) error {
	data := bson.M{
		"name": client.Name,
	}

	_, err := c.db.Collection("temp-client").InsertOne(ctx, data)

	return err
}

func (c client) Find(ctx context.Context, name string) (*models.Client, error) {
	data := bson.M{
		"name": name,
	}

	result := c.db.Collection("temp-client").FindOne(ctx, data)

	var client models.Client

	err := result.Decode(&client)
	if err != nil {
		return nil, err
	}

	return &client, nil
}
