package services

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/vltvdnl/birthday-notifier-rest/internal/domain/models"
)

type UserService struct {
	log  *slog.Logger
	repo UserRepo
	auth UserAuth
}

func New(s UserRepo, a UserAuth, log *slog.Logger) *UserService {
	return &UserService{repo: s, log: log, auth: a}
}
func (u *UserService) Register(ctx context.Context, user models.User, password string) (uid int64, err error) {
	const log_op = "UserService.Register"
	uid, err = u.auth.Register(ctx, user, password)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", log_op, err)
	}
	user.ID = uid
	err = u.repo.SaveUser(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("%s :%w", log_op, err)
	}
	return uid, nil
}
func (u *UserService) Login(ctx context.Context, email string, password string) (token string, err error) {
	const log_op = "UserService.Login"
	token, err = u.auth.Login(ctx, email, password)
	if err != nil {
		return "", fmt.Errorf("%s: %w", log_op, err)
	}
	return token, nil
}

func (u *UserService) GetUser(ctx context.Context, id int64) (*models.User, error) {
	user, err := u.repo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) SaveUser(ctx context.Context, user models.User) error {
	err := u.repo.SaveUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) NotificationChange(ctx context.Context, user_id int64) error {
	err := u.repo.NotificationChange(ctx, user_id)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserService) GetAllUsers(ctx context.Context) (*[]models.User, error) {
	users, err := u.repo.AllUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, err
}
func (u *UserService) Subscribe(ctx context.Context, follower int64, user int64) error {
	panic("not implemented")
}

func (u *UserService) Unsubscribe(ctx context.Context, follower int64, user int64) error {
	panic("not umpemented")
}
