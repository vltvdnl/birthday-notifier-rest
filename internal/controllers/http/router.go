package http

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/vltvdnl/birthday-notifier-rest/internal/services"
)

func NewRouter(handler *gin.Engine,
	log *slog.Logger,
	appSecret string,
	parser JWTParser,
	user services.UserUsecase) { // как то надо от этого избавится (appSecret и тд)
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	h1 := handler.Group("/auth")
	{
		newAuthRoutes(h1, user, log)
	}

	h2 := handler.Group("/")
	h2.Use(Authorize(log, parser))
	{
		newUserRoutes(h2, user, log)
	}

}
