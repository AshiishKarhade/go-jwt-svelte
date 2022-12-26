package main

import (
	"log"

	"github.com/AshiishKarhade/jwt-go-react/database"
	"github.com/AshiishKarhade/jwt-go-react/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()

	routes.Setup(app)

	log.Fatal(app.Listen(":3000"))
}
