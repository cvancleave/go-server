package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func NewToken(email string, userId int, key, issuer string, timeout int) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        strconv.Itoa(userId),
		Audience:  email,
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Duration(timeout) * time.Minute).Unix(),
	})

	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateToken(token, key, issuer string) error {

	// parse token
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return err
	}

	if !parsedToken.Valid {
		return fmt.Errorf("invalid token")
	}

	// get user claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("invalid claims")
	}

	// validate issuer
	iss, ok := claims["iss"].(string)
	if !ok || iss != issuer {
		return fmt.Errorf("invalid issuer")
	}

	// validate expiration unix datetime
	expiration, ok := claims["exp"].(float64)
	if !ok {
		return fmt.Errorf("invalid expiration")
	}

	if time.Now().After(time.Unix(int64(expiration), 0)) {
		return fmt.Errorf("expired token")
	}

	return nil
}

func GetTokenFromRequest(r *http.Request) (string, error) {

	tokenHeader, tokenOk := r.Header["Authorization"]
	if !tokenOk {
		return "", fmt.Errorf("missing auth header")
	}

	// request Authorization header should look like 'Bearer <jwt here>'
	return strings.Split(tokenHeader[0], "Bearer ")[1], nil
}
