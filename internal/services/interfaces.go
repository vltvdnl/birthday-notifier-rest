package services

import (
	"context"

	"github.com/vltvdnl/birthday-notifier-rest/internal/domain/models"
)

type UserRepo interface {
	GetUser(ctx context.Context, id int64) (*models.User, error)
	SaveUser(ctx context.Context, user models.User) error
	NotificationChange(ctx context.Context, user_id int64) error
	AllUsers(ctx context.Context) (*[]models.User, error)
	Subscribe(ctx context.Context, user int64, id int64) error
	Unsubscribe(ctx context.Context, user int64, id int64) error
	// Subscriptions(ctx context.Context, user models.User) (*[]models.User, error)
}
type UserUsecase interface {
	UserAuth
	GetUser(ctx context.Context, user_id int64) (*models.User, error)
	GetAllUsers(ctx context.Context) (*[]models.User, error)
	Subscribe(ctx context.Context, follower int64, user int64) error
	Unsubscribe(ctx context.Context, follower int64, user int64) error
	NotificationChange(ctx context.Context, user_id int64) error
}

type UserAuth interface {
	Register(ctx context.Context, user models.User, password string) (uid int64, err error)
	Login(ctx context.Context, email string, password string) (token string, err error)
}
