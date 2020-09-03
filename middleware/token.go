package middleware

import (
	"gitee.com/cristiane/micro-mall-api/vars"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	jwtSecret     = "&WJof0jaY4ByTHR2"
	jwtExpireTime = 2 * time.Hour
)

type Claims struct {
	UserName string `json:"user_name"`
	Uid      int    `json:"uid"`
	jwt.StandardClaims
}

func GenerateToken(username string, uid int) (string, error) {
	var expire = jwtExpireTime
	if vars.JwtSetting.TokenExpireSecond > 0 {
		expire = time.Duration(vars.JwtSetting.TokenExpireSecond) * time.Second
	}
	var secret = jwtSecret
	if vars.JwtSetting.Secret != "" {
		secret = vars.JwtSetting.Secret
	}
	nowTime := time.Now()
	expireTime := nowTime.Add(expire)

	claims := Claims{
		UserName: username,
		Uid:      uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "web_gin_template",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(secret))
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	var secret = jwtSecret
	if vars.JwtSetting.Secret != "" {
		secret = vars.JwtSetting.Secret
	}
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(secret), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
