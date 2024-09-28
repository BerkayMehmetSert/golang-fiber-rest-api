package repository

import (
	"context"
	"fmt"
	"golang-fiber-rest-api/internal/golang-fiber-rest-api/domain"
	"golang-fiber-rest-api/internal/golang-fiber-rest-api/pkg"
	"golang-fiber-rest-api/internal/golang-fiber-rest-api/pkg/utils"
)

// IUserRepository, kullanıcı deposu işlemleri için arayüzü tanımlar.
type IUserRepository interface {
	// Upsert, bir kullanıcı ekler veya günceller.
	Upsert(ctx context.Context, user *domain.User) error

	// GetById, belirli bir kimliğe sahip kullanıcıyı döndürür.
	GetById(ctx context.Context, id string) (*domain.User, error)

	// Get, tüm kullanıcıları döndürür.
	Get(ctx context.Context) ([]*domain.User, error)

	// GetByEmail, belirli bir e-posta adresine sahip kullanıcıyı döndürür.
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
}

// userRepository, IUserRepository arayüzünün bellek içi uygulamasıdır.
type userRepository struct {
	userList []*domain.User
}

// NewUserRepository, başlangıç verileri ile yeni bir userRepository örneği oluşturur.
func NewUserRepository() IUserRepository {
	return &userRepository{
		userList: utils.GetUserStub(),
	}
}

// Upsert, depoya bir kullanıcı ekler veya günceller.
func (r *userRepository) Upsert(ctx context.Context, user *domain.User) error {
	r.userList = append(r.userList, user)
	return nil
}

// GetById, belirli bir kimliğe sahip kullanıcıyı döndürür.
func (r *userRepository) GetById(ctx context.Context, id string) (*domain.User, error) {
	for _, user := range r.userList {
		if user.Id == id {
			return user, nil
		}
	}

	fmt.Printf(pkg.Msg.UserNotFoundByID, id)
	return nil, nil
}

// Get, depodaki tüm kullanıcıları döndürür.
func (r *userRepository) Get(ctx context.Context) ([]*domain.User, error) {
	users := r.userList

	if users == nil {
		fmt.Printf(pkg.Msg.NoUsersInDatabase)
		return make([]*domain.User, 0), nil
	}

	return r.userList, nil
}

// GetByEmail, belirli bir e-posta adresine sahip kullanıcıyı döndürür.
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	for _, user := range r.userList {
		if user.Email == email {
			return user, nil
		}
	}

	fmt.Printf(pkg.Msg.UserNotFoundByEmail, email)
	return nil, nil
}
