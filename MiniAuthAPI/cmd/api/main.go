package main

import (
	"MiniProject/internal/auth"
	"MiniProject/internal/handlers"
	"MiniProject/internal/jwt"
	"MiniProject/internal/middleware"
	"MiniProject/internal/password"
	"MiniProject/internal/repository"
	"net/http"
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

	authHandler := handlers.NewAuthHandeler(authService)

	profileHandler := handlers.NewProfileHander()

	protectedProfile :=
		middleware.AuthenticationMiddleware(
			http.HandlerFunc(profileHandler.Profile),
		)

	loggedProtectedProfile :=
		middleware.LoggingMiddleware(
			protectedProfile,
		)

	mux := http.NewServeMux()

	mux.HandleFunc("/login", authHandler.Login)

	mux.Handle("/profile", loggedProtectedProfile)

	http.ListenAndServe(":8080", mux)
}
