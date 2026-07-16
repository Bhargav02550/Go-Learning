package main

import (
	"MiniProject/auth"
	"MiniProject/jwt"
	"MiniProject/password"
	"MiniProject/repository"
	"fmt"
)

func main() {
	repo := repository.NewMemoryUserRepository()
	verifier := &password.PlainPasswordVerifier{}
	jwt := &jwt.SimpleJwtProvider{}

	authService := auth.NewAuthService(
		repo,
		verifier,
		jwt,
	)

	token, error := authService.Login("bhargav@gmail.com", "Bhargav@025")

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(token)
	}
}
