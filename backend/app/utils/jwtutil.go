package utils

import (
	"os"
	"time"

	"go_next_todo/domain/model"

	"github.com/golang-jwt/jwt/v5"
)

// MyCustomClaims struct
type MyCustomClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateJwtToken function
// Generate jwt token
// @param user model.User
// @return string, error
func GenerateJwtToken(user model.User) (string, error) {
	claims := MyCustomClaims{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "go_next_todo",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
