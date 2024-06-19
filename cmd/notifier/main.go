package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/vltvdnl/birthday-notifier-rest/internal/config"
	"github.com/vltvdnl/birthday-notifier-rest/internal/controllers/http"
	"github.com/vltvdnl/birthday-notifier-rest/internal/services"
	"github.com/vltvdnl/birthday-notifier-rest/internal/services/repo"
	"github.com/vltvdnl/birthday-notifier-rest/pkg/clients/sso/grpc"
	"github.com/vltvdnl/birthday-notifier-rest/pkg/httpserver"
	"github.com/vltvdnl/birthday-notifier-rest/pkg/jwt"
	"github.com/vltvdnl/birthday-notifier-rest/pkg/log"
	"github.com/vltvdnl/birthday-notifier-rest/pkg/storage"
)

const App_ID = 1

func main() {
	cfg := config.MustLoad()
	log := log.New(cfg.Env)
	log.Info("all started")
	fmt.Println(cfg)
	storage, err := storage.New(cfg.Storage_Url)
	repo := repo.New(log, *storage)
	if err != nil {
		panic("storage didnt inited")
	}
	auth, err := grpc.New(context.Background(), log, cfg.SSO.Address, App_ID, cfg.SSO.Timeout, cfg.SSO.RetriesCount)
	if err != nil {
		panic("sso isn't inited")
	}
	u := services.New(repo, auth, log)
	parser := jwt.New(cfg.AppSecret)
	router := gin.New()
	http.NewRouter(router, log, cfg.AppSecret, parser, u)
	httpServer := httpserver.New(router, cfg.HTTPServer.Address)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case <-interrupt:
		// l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		// l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Error("%w", err)
	}
}
