package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/guifgr/go-learn/ApiFiberJWT/database"
	"github.com/guifgr/go-learn/ApiFiberJWT/routes"
)

func main(){
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	err := app.Listen(":5000")
	if err != nil {
		panic("Api could not be created")
	}

}