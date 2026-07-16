package auth

import (
	"MiniProject/jwt"
	"MiniProject/password"
	"MiniProject/repository"
	"errors"
)

type AuthService struct {
	repo     *repository.MemoryUserRepository
	verifier password.PasswordVerifier
	jwt      jwt.JwtProvider
}

func NewAuthService(
	repo *repository.MemoryUserRepository,
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

	isValid := a.verifier.Verify(user.Password, password)

	if !isValid {
		return "", errors.New("Invalid Password")
	}

	token := a.jwt.GenerateToken(user)

	return token, nil
}
