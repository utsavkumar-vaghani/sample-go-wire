package repository

import (
	"context"

	"example.com/sample-go-wire/src/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type user struct {
	db *mongo.Database
}

func New(db *mongo.Database) UserRepository {
	return &user{
		db: db,
	}
}

func (u user) Add(ctx context.Context, user models.User) error {
	data := bson.M{
		"name": user.Name,
	}

	_, err := u.db.Collection("temp-user").InsertOne(ctx, data)

	return err
}

func (u user) Find(ctx context.Context, name string) (*models.User, error) {
	data := bson.M{
		"name": name,
	}

	result := u.db.Collection("temp-user").FindOne(ctx, data)

	var user models.User

	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
