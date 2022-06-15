package usecase

import "github.com/gofiber/fiber"

var (
	ErrUserAlreadyExists   = fiber.NewError(403, "user aleady exists")
	ErrClientAlreadyExists = fiber.NewError(403, "client aleady exists")
	ErrInvalidNameLength   = fiber.NewError(400, "maximum length of name 20 is allowed")
)
