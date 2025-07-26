package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/arsharaj/project-servo/config"
	"github.com/arsharaj/project-servo/logger"
	"github.com/arsharaj/project-servo/server"
	"github.com/rs/zerolog/log"
)

func main() {
	// Initialize structured logging
	logger.Init()
	// Initialize configuration management
	config.Load()

	srv := server.New(config.AppConfig.ServerAddress)

	// Start server in a goroutine
	go func() {
		log.Info().Msg("Starting server on " + config.AppConfig.ServerAddress)
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Server error")
		}
	}()

	// Listen for interrupt signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-stop // wait until signal is received
	log.Info().Msg("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), config.AppConfig.ShutdownGrace)
	defer cancel()

	if err := srv.ShutDown(&ctx); err != nil {
		log.Error().Err(err).Msg("Graceful shutdown failed")
	} else {
		log.Info().Msg("Server gracefully stopped")
	}
}
