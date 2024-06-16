package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	// "github.com/golang-jwt/jwt/v5"
)

type Parser struct{}

// TODO: сделать нормальнуюв валидацию данных из токена
func (p *Parser) Parse(tokenString string, appSecret string) (uid int64, email string, err error) {
	parsed, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(appSecret), nil
	})
	if err != nil {
		return 0, "", nil
	}
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", fmt.Errorf("unexpected error while parsing token")
	}
	uid = claims["uid"].(int64)
	email = claims["email"].(string)

	return uid, email, nil
}
