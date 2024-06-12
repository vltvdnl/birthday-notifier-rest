package grpc

import (
	"log/slog"

	ssov1 "github.com/vltvdnl/proto-contract/gen/go/sso"
)

type Client struct {
	api ssov1.AuthClient
	log *slog.Logger
}
