package http

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type JWTParser interface {
	Parse(tokenString string) (uid int64, email string, err error)
}

var (
	ErrInvalidToken = errors.New("invalid token")
)

func extractToken(c *gin.Context) string {
	authHead := c.GetHeader("Authorization")
	splitToken := strings.Split(authHead, "Bearer ")
	if len(splitToken) != 2 {
		return ""
	}
	return splitToken[1]
}

// TODO: can be some bugs
func Authorize(log *slog.Logger,
	parser JWTParser) gin.HandlerFunc {
	const log_op = "middleware.auth.New"

	log = log.With(slog.String("op", log_op))

	return func(c *gin.Context) {
		tokenStr := extractToken(c)
		if tokenStr == "" {
			log.Info("no token in reques")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No authorization"})
			return
		}
		uid, email, err := parser.Parse(tokenStr)
		if err != nil {
			log.Warn("failed to parse token", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No authorization"}) // TODO
			return
		}
		log.Info("user authorized", slog.Int64("uid:", uid), slog.String("email:", email))

		c.Set("uid", uid)
	}
}
