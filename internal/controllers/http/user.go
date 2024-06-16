package http

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vltvdnl/birthday-notifier-rest/internal/domain/models"
	"github.com/vltvdnl/birthday-notifier-rest/internal/services"
)

type userRoutes struct {
	u   services.UserUsecase
	log *slog.Logger
}

func newUserRoutes(handler *gin.RouterGroup, u services.UserUsecase, log *slog.Logger) {
	r := &userRoutes{u: u, log: log}
	h := handler.Group("/")
	{
		h.GET("/allusers", r.allusers)
		h.POST("/subscribe", r.subscribe)
		h.POST("unsubscribe", r.unsubscribe)
		h.GET("notifier-change", r.notifier)
	}
}

type allUsersResponse struct {
	Users []models.User `json:"users"`
}

func (r *userRoutes) allusers(c *gin.Context) {
	const log_op = "userRoutes.allusers"
	users, err := r.u.GetAllUsers(c.Request.Context())
	if err != nil {
		r.log.Error("%s: %w", log_op, err)
		errorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}
	c.JSON(http.StatusOK, allUsersResponse{Users: *users})
}

type subscribeRequest struct {
	User_id int64 `json:"user_id"`
}
type subscribeResponse struct {
	Text string `json:"status"`
}

func (r *userRoutes) subscribe(c *gin.Context) {
	const log_op = "userRoutes.subscribe"
	var req subscribeRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		r.log.Warn("%s: %w", log_op, err)
		errorResponse(c, http.StatusBadRequest, "invalid body")
		return
	}
	err := r.u.Subscribe(c.Request.Context(), c.GetInt64("uid"), req.User_id)
	if err != nil {
		r.log.Error("%s: %w", log_op, err)
		errorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}
	c.JSON(http.StatusOK, subscribeResponse{"success"})
}
func (r *userRoutes) unsubscribe(c *gin.Context) {
	const log_op = "userRoutes.unsubscribe"
	var req subscribeRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		r.log.Warn("%s: %w", log_op, err)
		errorResponse(c, http.StatusBadRequest, "invalid body")
		return
	}
	err := r.u.Unsubscribe(c.Request.Context(), c.GetInt64("uid"), req.User_id)
	if err != nil {
		r.log.Error("%s: %w", log_op, err)
		errorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}
	c.JSON(http.StatusOK, subscribeResponse{"success"})
}

type notifierResponse struct {
	Status bool `json:"status"`
}

func (r *userRoutes) notifier(c *gin.Context) {
	const log_op = "userRoutes.notifier"
	err := r.u.NotificationChange(c.Request.Context(), c.GetInt64("uid"))
	if err != nil {
		r.log.Error("%s: %w", log_op, err)
		errorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}
	c.JSON(http.StatusOK, notifierResponse{true})
}
