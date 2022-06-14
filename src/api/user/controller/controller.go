package controller

import (
	"example.com/sample-go-wire/src/api/models"
	"example.com/sample-go-wire/src/api/user/service"
	"github.com/gofiber/fiber/v2"
)

type user struct {
	service service.UserService
}

func New(service service.UserService) *user {
	return &user{
		service: service,
	}
}

func (u *user) Add(ctx *fiber.Ctx) error {
	var user models.User

	err := ctx.BodyParser(&user)
	if err != nil {
		return err
	}

	err = u.service.Add(ctx.Context(), user)
	if err != nil {
		return err
	}

	return nil
}
