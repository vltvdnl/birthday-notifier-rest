package grpc

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/vltvdnl/birthday-notifier-rest/internal/domain/models"
	ssov1 "github.com/vltvdnl/proto-contract/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	api ssov1.AuthClient
	log *slog.Logger
}

func New(ctx context.Context, log *slog.Logger, addr string, timeout time.Duration, retriesCount int) (*Client, error) {
	const log_op = "grpc.New"

	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retriesCount)),
		grpcretry.WithPerRetryTimeout(timeout),
	}
	logOpts := []grpclog.Option{
		grpclog.WithLogOnEvents(grpclog.PayloadReceived, grpclog.PayloadReceived),
	}
	con, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(grpclog.UnaryClientInterceptor(InterceptorLog(log), logOpts...),
			grpcretry.UnaryClientInterceptor(retryOpts...)))

	// cc, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()),
	// 	grpc.WithChainUnaryInterceptor(grpclog.UnaryClientInterceptor(InterceptorLog(log), logOpts...), grpcretry.UnaryClientInterceptor(retryOpts...)))

	if err != nil {
		return nil, fmt.Errorf("%s: %w", log_op, err)
	}

	grpcClient := ssov1.NewAuthClient(con)

	return &Client{
		api: grpcClient,
		log: log,
	}, nil
}

// InterceptorLogger adapts slog logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func InterceptorLog(l *slog.Logger) grpclog.Logger {
	return grpclog.LoggerFunc(func(ctx context.Context, lvl grpclog.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

func (c *Client) Register(ctx context.Context, user models.User, password string) (uid int64, err error) {

	const log_op = "gprc.Register"
	resp, err := c.api.Register(ctx, &ssov1.RegisterRequest{Email: user.Email, Password: password}) //TODO: как то зарефакторить, чтобы избежать зависимости от домена
	if err != nil {
		return 0, fmt.Errorf("%s: %w", log_op, err)
	}
	return resp.GetUserId(), nil
}
func (c *Client) Login(ctx context.Context, app_id int, email string, password string) (token string, err error) {
	const log_op = "grpc.Login"
	resp, err := c.api.Login(ctx, &ssov1.LoginRequest{AppId: int32(app_id), Email: email, Password: password})
	if err != nil {
		return "", fmt.Errorf("%s: %w", log_op, err)
	}
	return resp.GetToken(), nil
}
