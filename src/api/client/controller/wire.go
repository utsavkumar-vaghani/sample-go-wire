//+build wireinject

package controller

import (
	"example.com/sample-go-wire/src/api/client/repository"
	"example.com/sample-go-wire/src/api/client/service"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

func Wire(db *mongo.Database) *client {
	panic(wire.Build(New, service.New, repository.New))
}
