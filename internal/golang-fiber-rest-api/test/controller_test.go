package test

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang-fiber-rest-api/internal/golang-fiber-rest-api/application/controller"
	"golang-fiber-rest-api/internal/golang-fiber-rest-api/application/handler/user"
	"golang-fiber-rest-api/internal/golang-fiber-rest-api/domain"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockUserQueryService struct {
	mock.Mock
}

func (m *MockUserQueryService) GetById(ctx context.Context, id string) (*domain.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserQueryService) Get(ctx context.Context) ([]*domain.User, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.User), args.Error(1)
}

func (m *MockUserQueryService) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

type MockUserCommandHandler struct {
	mock.Mock
}

func (m *MockUserCommandHandler) Save(ctx context.Context, command user.Command) error {
	args := m.Called(ctx, command)
	return args.Error(0)
}

func TestSave_UserCreatedSuccessfully(t *testing.T) {
	mockUserQueryService := new(MockUserQueryService)
	mockUserCommandHandler := new(MockUserCommandHandler)
	controller := controller.NewUserController(mockUserQueryService, mockUserCommandHandler)

	// Set up the mock to expect the Save method call
	mockUserCommandHandler.On("Save", mock.Anything, mock.Anything).Return(nil)

	app := fiber.New()
	app.Post("/user", controller.Save)

	req := httptest.NewRequest("POST", "/user", strings.NewReader(`{"firstName":"John","lastName":"Doe","email":"john.doe@example.com","password":"password","age":30}`))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestSave_InvalidRequestBody(t *testing.T) {
	mockUserQueryService := new(MockUserQueryService)
	mockUserCommandHandler := new(MockUserCommandHandler)
	controller := controller.NewUserController(mockUserQueryService, mockUserCommandHandler)

	app := fiber.New()
	app.Post("/user", controller.Save)

	req := httptest.NewRequest("POST", "/user", strings.NewReader(`invalid json`))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestGetUserById_UserFound(t *testing.T) {
	mockUserQueryService := new(MockUserQueryService)
	mockUserCommandHandler := new(MockUserCommandHandler)
	controller := controller.NewUserController(mockUserQueryService, mockUserCommandHandler)

	user := &domain.User{Id: "1", FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"}
	mockUserQueryService.On("GetById", mock.Anything, "1").Return(user, nil)

	app := fiber.New()
	app.Get("/user/:userId", controller.GetUserById)

	req := httptest.NewRequest("GET", "/user/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetUserById_UserNotFound(t *testing.T) {
	mockUserQueryService := new(MockUserQueryService)
	mockUserCommandHandler := new(MockUserCommandHandler)
	controller := controller.NewUserController(mockUserQueryService, mockUserCommandHandler)

	mockUserQueryService.On("GetById", mock.Anything, "100").Return((*domain.User)(nil), errors.New("user not found"))

	app := fiber.New()
	app.Get("/user/:userId", controller.GetUserById)

	req := httptest.NewRequest("GET", "/user/100", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestGetUserById_UserIdEmpty(t *testing.T) {
	mockUserQueryService := new(MockUserQueryService)
	mockUserCommandHandler := new(MockUserCommandHandler)
	controller := controller.NewUserController(mockUserQueryService, mockUserCommandHandler)

	app := fiber.New()
	app.Get("/user/:userId", controller.GetUserById)

	req := httptest.NewRequest("GET", "/user/", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGetUser_UsersFound(t *testing.T) {
	mockUserQueryService := new(MockUserQueryService)
	mockUserCommandHandler := new(MockUserCommandHandler)
	controller := controller.NewUserController(mockUserQueryService, mockUserCommandHandler)

	users := []*domain.User{
		{Id: "1", FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
		{Id: "2", FirstName: "Jane", LastName: "Doe", Email: "jane.doe@example.com"},
	}
	mockUserQueryService.On("Get", mock.Anything).Return(users, nil)

	app := fiber.New()
	app.Get("/users", controller.GetUser)

	req := httptest.NewRequest("GET", "/users", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetUser_NoUsersFound(t *testing.T) {
	mockUserQueryService := new(MockUserQueryService)
	mockUserCommandHandler := new(MockUserCommandHandler)
	controller := controller.NewUserController(mockUserQueryService, mockUserCommandHandler)

	mockUserQueryService.On("Get", mock.Anything).Return(([]*domain.User)(nil), errors.New("no users found"))

	app := fiber.New()
	app.Get("/users", controller.GetUser)

	req := httptest.NewRequest("GET", "/users", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestGetUserByEmail_UserFound(t *testing.T) {
	mockUserQueryService := new(MockUserQueryService)
	mockUserCommandHandler := new(MockUserCommandHandler)
	controller := controller.NewUserController(mockUserQueryService, mockUserCommandHandler)

	user := &domain.User{Id: "1", FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"}
	mockUserQueryService.On("GetByEmail", mock.Anything, "john.doe@example.com").Return(user, nil)

	app := fiber.New()
	app.Get("/user/email/:email", controller.GetUserByEmail)

	req := httptest.NewRequest("GET", "/user/email/john.doe@example.com", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetUserByEmail_UserNotFound(t *testing.T) {
	mockUserQueryService := new(MockUserQueryService)
	mockUserCommandHandler := new(MockUserCommandHandler)
	controller := controller.NewUserController(mockUserQueryService, mockUserCommandHandler)

	mockUserQueryService.On("GetByEmail", mock.Anything, "john.doe@example.com").Return((*domain.User)(nil), errors.New("user not found"))

	app := fiber.New()
	app.Get("/user/email/:email", controller.GetUserByEmail)

	req := httptest.NewRequest("GET", "/user/email/john.doe@example.com", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestGetUserByEmail_EmailEmpty(t *testing.T) {
	mockUserQueryService := new(MockUserQueryService)
	mockUserCommandHandler := new(MockUserCommandHandler)
	controller := controller.NewUserController(mockUserQueryService, mockUserCommandHandler)

	app := fiber.New()
	app.Get("/user/email/:email", controller.GetUserByEmail)
	req := httptest.NewRequest("GET", "/user/email/", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}
