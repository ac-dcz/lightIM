package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type TokenOption struct {
	AccessSecret string
	AccessExpire int64
}

// BuildToken 生成token
func BuildToken(opt *TokenOption, payload map[string]interface{}) (string, error) {
	claims := make(jwt.MapClaims)
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Second * time.Duration(opt.AccessExpire)).Unix()
	for k, v := range payload {
		claims[k] = v
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(opt.AccessSecret))
}

// VerifyToken 验证Token
func VerifyToken(opt *TokenOption, tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(opt.AccessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	data := token.Claims.(jwt.MapClaims)
	payload := make(map[string]interface{})
	for k, v := range data {
		switch k {
		case "iat", "exp", "aud", "nbf":
		default:
			payload[k] = v
		}
	}
	return payload, nil
}
