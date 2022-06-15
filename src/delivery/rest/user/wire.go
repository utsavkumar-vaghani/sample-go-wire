//+build wireinject

package controller

import (
	repository "example.com/sample-go-wire/src/repository/user"
	usecase "example.com/sample-go-wire/src/usecase/user"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

func Wire(db *mongo.Database) *user {
	panic(wire.Build(New, usecase.New, repository.New))
}
