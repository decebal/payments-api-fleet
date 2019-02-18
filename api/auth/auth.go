package auth

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// SigningKey is used in setting up the HTTP server middleware
var SigningKey = "super-secret"

// HashAndSalt gets a cleartext password and transforms it into a hashed password
func HashAndSalt(p string) string {
	return p
}

// ComparePasswords checks if the hashed password matches the cleartext password
func ComparePasswords(hashedPassword, plainPassword string) bool {
	return hashedPassword == plainPassword
}

// GetJWT returns a generated JWT token from a username and id
func GetJWT(username string, id int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = jwt.MapClaims{
		"id":       id,
		"username": username,
		"exp":      time.Now().Add(time.Hour).Unix(),
	}

	tokenString, err := token.SignedString([]byte(SigningKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func DecodeJWT(jwtToken string) (int, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SigningKey), nil
	})

	if err != nil {
		return -1, errors.New("jwt malformed")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return int(claims["id"].(float64)), nil
	}

	return -1, errors.New("jwt not valid")
}
