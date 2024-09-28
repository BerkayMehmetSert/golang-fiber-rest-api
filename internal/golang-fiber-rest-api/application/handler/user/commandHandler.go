package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang-fiber-rest-api/internal/golang-fiber-rest-api/application/repository"
	"golang-fiber-rest-api/internal/golang-fiber-rest-api/domain"
	"golang-fiber-rest-api/internal/golang-fiber-rest-api/pkg"
)

type ICommandHandler interface {
	Save(ctx context.Context, command Command) error
}

// commandHandler, kullanıcı komutlarını işlemek için kullanılan yapıdır.
type commandHandler struct {
	userRepository repository.IUserRepository
}

// NewCommandHandler, yeni bir commandHandler örneği oluşturur.
func NewCommandHandler(userRepository repository.IUserRepository) ICommandHandler {
	return &commandHandler{userRepository: userRepository}
}

// Save, verilen komutu kullanarak bir kullanıcı kaydeder.
// Eğer kullanıcı zaten mevcutsa, bir hata döner.
func (h *commandHandler) Save(ctx context.Context, command Command) error {
	user, err := h.userRepository.GetByEmail(ctx, command.Email)

	if err != nil {
		return err
	}

	if user != nil {
		return errors.New(fmt.Sprintf(pkg.Msg.UserAlreadyExists, command.Email))
	}

	if err := h.userRepository.Upsert(ctx, h.BuildEntity(command)); err != nil {
		return err
	}

	return nil
}

// BuildEntity, verilen komutu kullanarak bir kullanıcı nesnesi oluşturur.
func (h *commandHandler) BuildEntity(command Command) *domain.User {
	return &domain.User{
		Id:        uuid.NewString(),
		FirstName: command.FirstName,
		LastName:  command.LastName,
		Email:     command.Email,
		Password:  command.Password,
		Age:       command.Age,
	}
}
