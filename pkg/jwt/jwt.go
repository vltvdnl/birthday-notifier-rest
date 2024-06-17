package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	// "github.com/golang-jwt/jwt/v5"
)

type Parser struct {
	AppSecret string
}

func New(appSecret string) *Parser {
	return &Parser{appSecret}
}

// TODO: сделать нормальнуюв валидацию данных из токена
func (p *Parser) Parse(tokenString string) (uid int64, email string, err error) {
	parsed, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(p.AppSecret), nil
	})
	if err != nil {
		return 0, "", fmt.Errorf("error while parsing: %w", err)
	}
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok || !parsed.Valid {
		return 0, "", fmt.Errorf("unexpected error while parsing token")
	}
	temp, ok := claims["uid"].(float64)
	if !ok {
		return 0, "", fmt.Errorf("pizdec")
	}
	uid = int64(temp)
	email = claims["email"].(string)

	return uid, email, nil
}
