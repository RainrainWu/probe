package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/RainrainWu/probe/pkg/config"
)

type UserData struct {
	username string	`json:"username"`
	password string	`json:"password"`
}

type Claims struct {
	role string		`json:"role"`
	jwt.StandardClaims
}

var (
	jwtSecret = []byte(config.JWT_SECRET)
)

func CheckUser(user UserData) bool {
	
	// user check
	u_check := user.username == config.USERNAME
	p_check := user.password == config.PASSWORD
	if u_check && p_check {
		return true
	}
	return false
}

func GenToken(user UserData) string {

	// generate jwt
	now := time.Now()
	claims := Claims{
		role:           "Tester",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(20 * time.Second).Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    "ginJWT",
		},
	}
	
	// sign token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		return ""
	}
	return token
}

func CheckToken(token string) *Claims {

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})
	if err != nil {
		var message string
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors & jwt.ValidationErrorMalformed != 0 {
				message = "token is malformed"
			} else if ve.Errors & jwt.ValidationErrorUnverifiable != 0{
				message = "token could not be verified because of signing problems"
			} else if ve.Errors & jwt.ValidationErrorSignatureInvalid != 0 {
				message = "signature validation failed"
			} else if ve.Errors & jwt.ValidationErrorExpired != 0 {
				message = "token is expired"
			} else if ve.Errors & jwt.ValidationErrorNotValidYet != 0 {
				message = "token is not yet valid before sometime"
			} else {
				message = "can not handle this token"
			}
			fmt.Println(message)
		}
		return 
		


		
	}

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		fmt.Println("role:", claims.role)
		return claims
	}
	return nil
}

func main() {

}
