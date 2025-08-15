package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	// Define the secret key used to sign the token
	secretKey := []byte("PraveenHegde")

	// Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = "praveen.hegde"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Failed to sign the token:", err)
		return
	}

	// Print the token
	fmt.Println("JWT Token:", tokenString)
}
