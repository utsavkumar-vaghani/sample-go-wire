package controller

import (
	domain "example.com/sample-go-wire/src/domain/user"
	"github.com/gofiber/fiber/v2"
)

type user struct {
	domain domain.UserService
}

func New(domain domain.UserService) *user {
	return &user{
		domain: domain,
	}
}

func (u *user) Add(ctx *fiber.Ctx) error {
	var user domain.User

	err := ctx.BodyParser(&user)
	if err != nil {
		return err
	}

	err = u.domain.Add(ctx.Context(), user)
	if err != nil {
		return err
	}

	return nil
}
