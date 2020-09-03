package main

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	SignKey = "xxxx"
)

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// CustomClaims .
type CustomClaims struct {
	UID   int32  `json:"userId"`
	Phone string `json:"name"`
	jwt.StandardClaims
}

// JWT .
type JWT struct {
	SigningKey []byte
}

// NewJWT .
func NewJWT() *JWT {
	return &JWT{
		[]byte(SignKey),
	}
}

func (j *JWT) CreateToken(uid int32, phone string, expire int64) (string, error) {
	claims := CustomClaims{
		UID:   uid,
		Phone: phone,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(expire)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken .
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}
