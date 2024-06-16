package main

import (
	"fmt"
	"log"

	"github.com/vltvdnl/birthday-notifier-rest/internal/config"
)

func main() {
	cfg := config.MustLoad()
	log := log.New(cfg.Env)
	log.Info("all started")
	fmt.Println(cfg)
}
