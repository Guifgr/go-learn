package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guifgr/go-learn/ApiFiberJWT/controller"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Post("/api/login/cookie", controller.LoginCookie)
	app.Get("/api/user", controller.User)
	app.Post("/api/logout", controller.Logout)
	app.Post("/api/forgot", controller.Forgot)
	app.Post("/api/reset", controller.Reset)
}
