package web

import (
	"github.com/gofiber/fiber/v2"
	"golang-fiber-rest-api/internal/golang-fiber-rest-api/application/controller"
	"net/http"
)

func InitRouter(app *fiber.App, userController controller.IUserController) {

	app.Get("/healthcheck", func(context *fiber.Ctx) error {
		return context.SendStatus(http.StatusOK)
	})

	routeGroup := app.Group("/api/v1")

	routeGroup.Get("/user", userController.GetUser)
	routeGroup.Post("/user", userController.Save)
	routeGroup.Get("/user/:userId", userController.GetUserById)
	routeGroup.Get("/user/email/:email", userController.GetUserByEmail)
}
