package jwt

import "MiniProject/models"

type JwtProvider interface {
	GenerateToken(user *models.User) string
}

type SimpleJwtProvider struct{}

func (j *SimpleJwtProvider) GenerateToken(user *models.User) string {
	return "JWT_TOKEN_FOR_" + user.Email
}
