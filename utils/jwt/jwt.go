package jwt

import (
	"FIFA/core"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TokenExpireDuration = time.Hour * 2

type JWT struct {
	SigningKey []byte
}

type CustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

var (
	TokenExpired     = errors.New("Token已过期")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("还未生成token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(core.Conf.SigningKey),
	}
}

// GenToken 生成Token
func (j *JWT) GenToken(name string) (AuthToken string, err error) {
	claims := CustomClaims{
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "FIFA",                                     // 签发人
		},
	}
	AuthToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.SigningKey)
	return
}

// ParseToken 解析Token
func (j *JWT) ParseToken(token string) (*CustomClaims, error) {
	tokenString, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			return nil, TokenMalformed
		} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
			//token已经到期
			return nil, TokenExpired
		} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
			//token还未激活
			return nil, TokenNotValidYet
		} else {
			return nil, TokenInvalid
		}
	}
	if tokenString != nil {
		if claims, ok := tokenString.Claims.(*CustomClaims); ok && tokenString.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

// RefreshToken 刷新Token
func (j *JWT) RefreshToken(token string) (newToken string, err error) {
	//从旧的token中解析claims数据
	var claims CustomClaims
	_, err = jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	//拿到的claims重新生成token
	return j.GenToken(claims.Name)
}
