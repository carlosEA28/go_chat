package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/carlosEA28/go_chat/internal/config"
	"github.com/carlosEA28/go_chat/internal/logger"
	"github.com/carlosEA28/go_chat/internal/repositories"
	"github.com/carlosEA28/go_chat/internal/server"
)

func main() {
	log := logger.New()
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	mongoCtx, err := repositories.NewMongoRepositoryContext(cfg.Mongo.URI, cfg.Mongo.Database, cfg.Mongo.Collection)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to mongo")
	}
	log.Info().Str("database", cfg.Mongo.Database).Msg("connected to mongo")

	srv := server.New(cfg, &log)

	router := srv.SetupRoutes()

	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Server.Port),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Info().Str("port", cfg.Server.Port).Msg("starting http server")
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("failed to start http server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msg("shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("failed to shutdown http server")
		return
	}

	log.Info().Msg("shutting down database")
	if err := mongoCtx.Close(ctx); err != nil {
		log.Error().Err(err).Msg("failed to disconnect mongo")
		return
	}

}
