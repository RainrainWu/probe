package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/RainrainWu/probe/pkg/config"
)

type userData struct {
	Username string	`json:"username"`
	Password string	`json:"password"`
}

type Claims struct {
	User string		`json:"user"`
	Role string		`json:"role"`
	jwt.StandardClaims
}

var (
	jwtSecret = []byte(config.JWT_SECRET)
)

func NewUser() *userData {
	
	return &userData{}
}

func (user *userData) ShowUser() {

	fmt.Println(user.Username)
	fmt.Println(user.Password)
}

func (user *userData) CheckUser() bool {
	
	// user check
	u_check := user.Username == config.USERNAME
	p_check := user.Password == config.PASSWORD
	if u_check && p_check {
		return true
	}
	return false
}

func (user *userData) GenToken() string {

	// generate jwt
	now := time.Now()
	claims := Claims{
		User:			user.Username,
		Role:           "Tester",
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

func ValidateToken(token string) (string, *jwt.Token) {

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
		}
		return message, nil
	} else {
		return "", tokenClaims
	}
}
