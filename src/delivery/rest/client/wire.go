//+build wireinject

package controller

import (
	repository "example.com/sample-go-wire/src/repository/client"
	usecase "example.com/sample-go-wire/src/usecase/client"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

func Wire(db *mongo.Database) *client {
	panic(wire.Build(New, usecase.New, repository.New))
}
