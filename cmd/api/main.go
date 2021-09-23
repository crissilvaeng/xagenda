package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/crissilvaeng/xagenda/internal/app/api"
	"github.com/crissilvaeng/xagenda/internal/pkg/support"
	"github.com/crissilvaeng/xagenda/pkg/people"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	var cfg support.Config
	err := cfg.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err.Error())
	}

	logger := support.NewLogger(support.LogLevel(cfg.LogLevel))

	db, err := gorm.Open(postgres.Open(cfg.Dsn), &gorm.Config{})
	if err != nil {
		logger.Error("failed to connect database: %v", err.Error())
		os.Exit(2)
	}

	if err := db.AutoMigrate(&people.PersonInfo{}); err != nil {
		logger.Errorf("failed to apply migrations: %v", err.Error())
		os.Exit(2)
	}

	r := mux.NewRouter()
	r.Handle("/metrics", promhttp.Handler())

	s := r.PathPrefix("/api/v1").Subrouter()
	api.NewService(s, db, logger)

	handler := handlers.CombinedLoggingHandler(os.Stdout, r)
	handler = http.TimeoutHandler(handler, cfg.Timeout, "timeout")

	srv := &http.Server{
		Addr:         cfg.Addr,
		Handler:      handlers.RecoveryHandler()(handler),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		ErrorLog:     logger.LogError,
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.Error(err)
		os.Exit(2)
	}
}
