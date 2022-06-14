package controller

import (
	"example.com/sample-go-wire/src/api/client/service"
	"example.com/sample-go-wire/src/api/models"
	"github.com/gofiber/fiber/v2"
)

type client struct {
	service service.ClientService
}

func New(service service.ClientService) *client {
	return &client{
		service: service,
	}
}

func (c *client) Add(ctx *fiber.Ctx) error {
	var client models.Client

	err := ctx.BodyParser(&client)
	if err != nil {
		return err
	}

	err = c.service.Add(ctx.Context(), client)
	if err != nil {
		return err
	}

	return nil
}
