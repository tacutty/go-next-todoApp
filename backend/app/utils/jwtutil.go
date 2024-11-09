package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go_next_todo/domain/model"
)

// mySigningKey variable
var mySigningKey = []byte(os.Getenv("SECRET"))

// MyCustomClaims struct
type MyCustomClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// GenerateJwtToken function
// Generate jwt token
// @param user model.User
// @return string, error
func GenerateJwtToken(user model.User) (string, error) {
	claims := MyCustomClaims{
		user.ID,
		user.Username,
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "go_next_todo",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
