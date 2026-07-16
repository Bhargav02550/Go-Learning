package repository

import (
	"MiniProject/internal/models"
	"errors"
)

type UserRepository interface {
	GetByEmail(email string) (*models.User, error)
}

type MemoryUserRepository struct {
	Users map[string]*models.User
}

func NewMemoryUserRepository() *MemoryUserRepository {
	users := make(map[string]*models.User)

	users["bhargav@gmail.com"] = &models.User{
		Id:       1,
		Email:    "bhargav@gmail.com",
		Password: "Bhargav@025",
	}

	users["pavan@gmail.com"] = &models.User{
		Id:       2,
		Email:    "pavan@gmail.com",
		Password: "Pavan@123",
	}

	return &MemoryUserRepository{
		Users: users,
	}
}

func (m *MemoryUserRepository) GetByEmail(email string) (*models.User, error) {
	user, found := m.Users[email]

	if !found {
		return nil, errors.New("User Not Found")
	}

	return user, nil
}
