package controller

import (
	domain "example.com/sample-go-wire/src/domain/client"
	"github.com/gofiber/fiber/v2"
)

type client struct {
	domain domain.ClientService
}

func New(domain domain.ClientService) *client {
	return &client{
		domain: domain,
	}
}

func (c *client) Add(ctx *fiber.Ctx) error {
	var client domain.Client

	err := ctx.BodyParser(&client)
	if err != nil {
		return err
	}

	err = c.domain.Add(ctx.Context(), client)
	if err != nil {
		return err
	}

	return nil
}
