package middleware

import (
	"example.com/sample-go-wire/src/logger"
	"github.com/gofiber/fiber/v2"
)

func Logging(ctx *fiber.Ctx) error {
	var method = string(ctx.Request().Header.Method())
	
	err := ctx.Next()
	if err != nil {
		code := 500

		if val, ok := err.(*fiber.Error); ok {
			code = val.Code
		}

		logger.GetLogger().With("method", method, "status-code", code).Info()
		logger.GetLogger().With("method", method, "status-code", code).Error(err)
	} else {
		logger.GetLogger().With("method", method, "status-code", ctx.Response().StatusCode()).Info()
	}

	return err
}
