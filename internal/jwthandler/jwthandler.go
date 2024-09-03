package jwthandler

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtHandler struct {
	Secret string
}

func GetJwt() *JwtHandler {
	return &JwtHandler{}
}

func (obj *JwtHandler) GetToken(lifepan time.Duration, id int) (token string, err error) {
	exp := time.Now().Add(lifepan)
	claims := jwt.MapClaims{
		"exp":     exp.Unix(),
		"iat":     time.Now().Unix(),
		"nbf":     time.Now().Unix(),
		"user_id": id,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwtToken.SignedString([]byte(obj.Secret))
	return
}
