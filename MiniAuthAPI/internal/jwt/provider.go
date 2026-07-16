package jwt

import "MiniProject/internal/models"

type JwtProvider interface {
	GenerateToken(user *models.User) string
}

type SimpleJwtProvider struct{}

func (p *SimpleJwtProvider) GenerateToken(user *models.User) string {
	return "JWT_TOKEN_FOR_" + user.Email
}
