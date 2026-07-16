package auth

import (
	"MiniProject/internal/jwt"
	"MiniProject/internal/password"
	"MiniProject/internal/repository"
	"errors"
)

type AuthService struct {
	repo     repository.UserRepository
	verifier password.PasswordVerifier
	jwt      jwt.JwtProvider
}

func NewAuthService(
	repo repository.UserRepository,
	verifier password.PasswordVerifier,
	jwt jwt.JwtProvider,
) *AuthService {
	return &AuthService{
		repo:     repo,
		verifier: verifier,
		jwt:      jwt,
	}
}

func (a *AuthService) Login(email string, password string) (string, error) {
	user, err := a.repo.GetByEmail(email)

	if err != nil {
		return "", err
	}

	isValid := a.verifier.VerifyPassword(user.Password, password)

	if !isValid {
		return "", errors.New("Invalid Password")
	}

	token := a.jwt.GenerateToken(user)

	return token, nil
}
