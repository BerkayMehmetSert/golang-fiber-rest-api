package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"golang-fiber-rest-api/internal/golang-fiber-rest-api/application/controller"
	"golang-fiber-rest-api/internal/golang-fiber-rest-api/application/handler/user"
	"golang-fiber-rest-api/internal/golang-fiber-rest-api/application/query"
	"golang-fiber-rest-api/internal/golang-fiber-rest-api/application/repository"
	"golang-fiber-rest-api/internal/golang-fiber-rest-api/application/web"
	"golang-fiber-rest-api/internal/golang-fiber-rest-api/pkg/server"
)

func main() {
	// fiber framework http server
	app := fiber.New()

	app.Use(recover.New())

	userRepository := repository.NewUserRepository()
	userQueryService := query.NewUserQueryService(userRepository)
	userCommandHandler := user.NewCommandHandler(userRepository)
	userController := controller.NewUserController(userQueryService, userCommandHandler)

	// Router
	web.InitRouter(app, userController)

	// Server
	server.NewServer(app).StartHttpServer()
}
