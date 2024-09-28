package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang-fiber-rest-api/internal/golang-fiber-rest-api/application/controller/request"
	"golang-fiber-rest-api/internal/golang-fiber-rest-api/application/controller/response"
	"golang-fiber-rest-api/internal/golang-fiber-rest-api/application/handler/user"
	"golang-fiber-rest-api/internal/golang-fiber-rest-api/application/query"
	"golang-fiber-rest-api/internal/golang-fiber-rest-api/pkg"
	"net/http"
)

// IUserController, kullanıcı kontrol işlemleri için arayüzü tanımlar.
type IUserController interface {
	// Save, yeni bir kullanıcı kaydeder.
	Save(ctx *fiber.Ctx) error

	// GetUserById, belirli bir kimliğe sahip kullanıcıyı döndürür.
	GetUserById(ctx *fiber.Ctx) error

	// GetUser, tüm kullanıcıları döndürür.
	GetUser(ctx *fiber.Ctx) error

	// GetUserByEmail, belirli bir e-posta adresine sahip kullanıcıyı döndürür.
	GetUserByEmail(ctx *fiber.Ctx) error
}

// userController, IUserController arayüzünün uygulamasıdır.
type userController struct {
	userQueryService   query.IUserQueryService
	userCommandHandler user.ICommandHandler
}

// NewUserController, yeni bir userController örneği oluşturur.
func NewUserController(userQueryService query.IUserQueryService, userCommandHandler user.ICommandHandler) IUserController {
	return &userController{
		userQueryService:   userQueryService,
		userCommandHandler: userCommandHandler,
	}
}

// Save, yeni bir kullanıcı kaydeder.
func (c *userController) Save(ctx *fiber.Ctx) error {
	var req request.UserCreteRequest
	err := ctx.BodyParser(&req)

	if err != nil {
		fmt.Printf(pkg.Msg.JsonBindingError, err.Error())
		return err
	}

	fmt.Printf(pkg.Msg.SaveStarted, req)

	if err = c.userCommandHandler.Save(ctx.UserContext(), req.ToCommand()); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(pkg.Msg.UserCreatedSuccess)
}

// GetUserById, belirli bir kimliğe sahip kullanıcıyı döndürür.
func (c *userController) GetUserById(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")

	if userId == "" {
		return ctx.Status(http.StatusBadRequest).JSON(pkg.Msg.UserIdCannotBeEmpty)
	}

	fmt.Printf(pkg.Msg.GetUserByIdStarted, userId)

	user, err := c.userQueryService.GetById(ctx.UserContext(), userId)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(response.ToUserResponse(user))
}

// GetUser, tüm kullanıcıları döndürür.
func (c *userController) GetUser(ctx *fiber.Ctx) error {
	fmt.Printf(pkg.Msg.GetUserStarted)

	users, err := c.userQueryService.Get(ctx.UserContext())

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(response.ToUserResponseList(users))
}

// GetUserByEmail, belirli bir e-posta adresine sahip kullanıcıyı döndürür.
func (c *userController) GetUserByEmail(ctx *fiber.Ctx) error {
	email := ctx.Params("email")

	if email == "" {
		return ctx.Status(http.StatusBadRequest).JSON(pkg.Msg.EmailCannotBeEmpty)
	}

	fmt.Printf(pkg.Msg.GetUserByEmailStarted, email)

	user, err := c.userQueryService.GetByEmail(ctx.UserContext(), email)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(response.ToUserResponse(user))
}
