package repository

import (
	"MiniProject/models"
	"errors"
)

type MemoryUserRepository struct {
	Users map[string]models.User
}

func NewMemoryUserRepository() *MemoryUserRepository {
	users := make(map[string]models.User)

	users["bhargav@gmail.com"] = models.User{
		Id:       1,
		Email:    "bhargav@gmail.com",
		Password: "Bhargav@025",
	}

	users["pavan@gmail.com"] = models.User{
		Id:       2,
		Email:    "pavan@gmail.com",
		Password: "Hello123",
	}

	return &MemoryUserRepository{
		Users: users,
	}
}

func (m *MemoryUserRepository) GetByEmail(email string) (*models.User, error) {
	user, found := m.Users[email]

	if !found {
		return nil, errors.New("User not found")
	}
	return &user, nil
}
