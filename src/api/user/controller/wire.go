//+build wireinject

package controller

import (
	"example.com/sample-go-wire/src/api/user/repository"
	"example.com/sample-go-wire/src/api/user/service"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

func Wire(db *mongo.Database) *user {
	panic(wire.Build(New, service.New, repository.New))
}
