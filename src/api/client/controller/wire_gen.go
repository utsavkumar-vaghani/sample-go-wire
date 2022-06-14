// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package controller

import (
	"example.com/sample-go-wire/src/api/client/repository"
	"example.com/sample-go-wire/src/api/client/service"
	"go.mongodb.org/mongo-driver/mongo"
)

// Injectors from wire.go:

func Wire(db *mongo.Database) *client {
	clientRepository := repository.New(db)
	clientService := service.New(clientRepository)
	controllerClient := New(clientService)
	return controllerClient
}