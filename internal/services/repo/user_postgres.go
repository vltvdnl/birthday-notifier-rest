package repo

import (
	"context"
	"log/slog"

	"github.com/vltvdnl/birthday-notifier-rest/internal/domain/models"
	"github.com/vltvdnl/birthday-notifier-rest/pkg/storage"
)

type UserRepo struct {
	log *slog.Logger
	s   storage.Storage
}

func (r *UserRepo) GetUser(ctx context.Context, user_id int64) (*models.User, error) {
	const log_op = "UserRepo.GetUser"
	r.log = r.log.With(slog.String("op", log_op))
	sqlstatement := `SELECT id, first_name, last_name, email, birthdate FROM users WHERE id = $1`
	var user models.User
	err := r.s.QueryRowContext(ctx, sqlstatement, user).Scan(&user.ID, &user.First_name, &user.Last_name, &user.Email, &user.Birthdate)
	if err != nil {
		r.log.Error("can't get user from db: ", err)
	}
	return &user, nil

}
func (r *UserRepo) SaveUser(ctx context.Context, user models.User) error {
	const log_op = "UserRepo.SaveUser"
	r.log = r.log.With(slog.String("op", log_op))

	sqlstatement := `INSERT INTO users(id, first_name, last_name, email, birthdate) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.s.ExecContext(ctx, sqlstatement, user.ID, user.First_name, user.Last_name, user.Email, user.Birthdate)
	if err != nil {
		r.log.Error("can't save user to db ", err)
		return err
	}
	return nil
}
func (r *UserRepo) NotificationChange(ctx context.Context, user_id int64) error {
	const log_op = "UserRepo.NotificationChange"
	r.log = r.log.With(slog.String("op", log_op))

	sqlstatement := `UPDATE users SET want_notifications = NOT want_notifications WHERE user_id = $1`
	_, err := r.s.ExecContext(ctx, sqlstatement, user_id)
	if err != nil {
		r.log.Error("can't update notifications status ", err)
		return err
	}
	return nil
}
func (r *UserRepo) AllUsers(ctx context.Context) (*[]models.User, error) {
	const log_op = "UserRepo.AllUsers"

	r.log = r.log.With(slog.String("op", log_op))

	sqlstatement := `SELECT id, first_name, last_name, email, birthdate FROM users`
	users := make([]models.User, 0, 150)

	rows, err := r.s.QueryContext(ctx, sqlstatement)
	if err != nil {
		r.log.Error("can't scan rows from db", err)
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.First_name, &user.Last_name, &user.Email, &user.Birthdate)
		if err != nil {
			r.log.Error("can't read row", err)
			return nil, err
		}
		users = append(users, user)
	}
	return &users, nil

}
func (r *UserRepo) Subscribe(ctx context.Context, follower_id int64, user int64) error {
	panic("not impelented")
}
func (r *UserRepo) Unsubscribe(ctx context.Context, follower_id int64, user int64) error {
	panic("not impemented")
}
func (r *UserRepo) Subscriptions(ctx context.Context, user_id int64) (*[]models.User, error) {
	panic("not implemented")
}
