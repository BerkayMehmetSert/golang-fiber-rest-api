package query

import (
	"context"
	"errors"
	"golang-fiber-rest-api/internal/golang-fiber-rest-api/application/repository"
	"golang-fiber-rest-api/internal/golang-fiber-rest-api/domain"
	"golang-fiber-rest-api/internal/golang-fiber-rest-api/pkg"
)

// IUserQueryService, kullanıcı sorgu işlemleri için arayüzü tanımlar.
type IUserQueryService interface {
	// GetById, belirli bir kimliğe sahip kullanıcıyı döndürür.
	GetById(ctx context.Context, id string) (*domain.User, error)

	// Get, tüm kullanıcıları döndürür.
	Get(ctx context.Context) ([]*domain.User, error)

	// GetByEmail, belirli bir e-posta adresine sahip kullanıcıyı döndürür.
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
}

// userQueryService, IUserQueryService arayüzünün uygulamasıdır.
type userQueryService struct {
	userRepository repository.IUserRepository
}

// NewUserQueryService, yeni bir userQueryService örneği oluşturur.
func NewUserQueryService(userRepository repository.IUserRepository) IUserQueryService {
	return &userQueryService{
		userRepository: userRepository,
	}
}

// GetById, belirli bir kimliğe sahip kullanıcıyı döndürür.
func (q *userQueryService) GetById(ctx context.Context, id string) (*domain.User, error) {
	user, err := q.userRepository.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New(pkg.Msg.NotFoundError)
	}

	return user, nil
}

// Get, tüm kullanıcıları döndürür.
func (q *userQueryService) Get(ctx context.Context) ([]*domain.User, error) {
	users, err := q.userRepository.Get(ctx)

	if err != nil {
		return nil, err
	}

	if users == nil {
		return nil, errors.New(pkg.Msg.NotFoundUsers)
	}

	return users, nil
}

// GetByEmail, belirli bir e-posta adresine sahip kullanıcıyı döndürür.
func (q *userQueryService) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := q.userRepository.GetByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New(pkg.Msg.NotFoundError)
	}

	return user, nil
}
