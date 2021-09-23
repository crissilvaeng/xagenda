package main

import (
	"log"

	"github.com/crissilvaeng/xagenda/internal/pkg/support"
)

func main() {
	var cfg support.Config
	err := cfg.Load()
	if err != nil {
		log.Fatal(cfg)
	}

	logger := support.NewLogger(support.LogLevel(cfg.LogLevel))
	logger.Info(cfg)

}
