package main

import (
	"log"

	"github.com/AshiishKarhade/jwt-go-react/database"
	"github.com/AshiishKarhade/jwt-go-react/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	app := fiber.New()

	//need for authentication
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	log.Fatal(app.Listen(":3000"))
}
