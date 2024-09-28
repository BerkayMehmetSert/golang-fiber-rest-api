package utils

import "golang-fiber-rest-api/internal/golang-fiber-rest-api/domain"

// GetUserStub, örnek kullanıcı verilerini döndüren bir fonksiyondur.
func GetUserStub() []*domain.User {
	return []*domain.User{
		{
			Id:        "1",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@gmail.com",
			Password:  "1234",
			Age:       26,
		},
		{
			Id:        "2",
			FirstName: "Jack",
			LastName:  "Doe",
			Email:     "jack@gmail.com",
			Password:  "123456",
			Age:       35,
		},
		{
			Id:        "3",
			FirstName: "Jane",
			LastName:  "Doe",
			Email:     "jane@gmail.com",
			Password:  "123456",
			Age:       11,
		},
	}
}
