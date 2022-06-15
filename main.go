package main

import (
	"log"

	// loading configs
	_ "example.com/sample-go-wire/src/config"

	"example.com/sample-go-wire/src/db"
	controllerClient "example.com/sample-go-wire/src/delivery/rest/client"
	controllerUser "example.com/sample-go-wire/src/delivery/rest/client"
	middleware "example.com/sample-go-wire/src/middlewares"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(middleware.Logging)

	db := db.NewMongoConnection()

	userController := controllerUser.Wire(db)
	clientController := controllerClient.Wire(db)

	app.Post("/user", userController.Add)
	app.Post("/client", clientController.Add)

	log.Fatal(app.Listen(":8080"))
}
