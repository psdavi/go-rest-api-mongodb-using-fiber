package main

import (
	"github.com/gofiber/fiber/v2"
	"goRestApiMongoDBFiber/configs"
	"goRestApiMongoDBFiber/routes"
)

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(app) //add this

	app.Listen(":6000")
}
