package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guifgr/go-learn/ApiFiberJWT/controller"
)

func Setup(app *fiber.App){
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Get("/api/user", controller.User)
	app.Get("/api/logout", controller.Logout)
}
