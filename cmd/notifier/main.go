package main

import (
	"fmt"
	"notifier/internal/config"
	"notifier/pkg/log"
)

func main() {
	cfg := config.MustLoad()
	log := log.New(cfg.Env)
	log.Info("all started")
	fmt.Println(cfg)
}
