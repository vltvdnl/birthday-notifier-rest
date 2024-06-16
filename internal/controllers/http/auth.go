package http

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vltvdnl/birthday-notifier-rest/internal/domain/models"
	"github.com/vltvdnl/birthday-notifier-rest/internal/services"
)

type authRoutes struct {
	u   services.UserUsecase
	log *slog.Logger
}

func newAuthRoutes(handler *gin.RouterGroup, u services.UserUsecase, log *slog.Logger) {
	r := &authRoutes{u, log}
	h := handler.Group("/")
	{
		h.POST("register", r.register)
		h.POST(("login"), r.login)
	}
}

type authRequest struct {
	First_name string    `json:"first_name"`
	Last_name  string    `json:"last_name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Birthdate  time.Time `json:"birthdate"`
}
type authResponse struct {
	UID int64 `json:"uid"`
}

func (r *authRoutes) register(c *gin.Context) {
	const log_op = "authRoutes.register"
	var request authRequest

	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		r.log.Error("%s: %w", log_op, err)
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	uid, err := r.u.Register(c.Request.Context(), models.User{
		First_name: request.First_name,
		Last_name:  request.Last_name,
		Email:      request.Email,
		Birthdate:  request.Birthdate,
	}, request.Password)
	if err != nil { //TODO: сделать валидацию того, что пользователь уже зарегистрирован
		r.log.Warn("%s: %w", log_op, err)
		errorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}
	c.JSON(http.StatusOK, authResponse{UID: uid})
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type loginResponse struct {
	Token string `json:"token"`
}

func (r *authRoutes) login(c *gin.Context) {
	const log_op = "authRoutes.login"
	var request loginRequest
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		r.log.Error("%s: %w", log_op, err)
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	token, err := r.u.Login(c.Request.Context(), request.Email, request.Password)
	if err != nil {
		r.log.Error("%s: %w", log_op, err)
		errorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}
	c.JSON(http.StatusOK, loginResponse{token})
}
