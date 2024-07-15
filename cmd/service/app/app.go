package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/ksusonic/goshorter/internal/repository"
	"github.com/ksusonic/goshorter/internal/service/shortener"
)

const (
	readTimeout  = 5 * time.Second
	WriteTimeout = 10 * time.Second
	stopTimeout  = time.Second * 5
)

func Run() {
	var (
		ctx = context.Background()
		cfg = newConfig()
		log = newLogger()
	)

	repo, closer := repository.NewRepository(ctx, cfg.DatabaseDSN)
	defer closer()

	shortenerService := shortener.NewService(cfg.ShortURLPrefix, repo, log)

	r := setupRouter(log, shortenerService)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      r.Handler(),
		ReadTimeout:  readTimeout,
		WriteTimeout: WriteTimeout,
	}

	log.Debug("starting server", zap.Int("port", cfg.Port))

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("listen and serve", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Info("shutting down server...", zap.Duration("timeout", stopTimeout))

	ctx, cancel := context.WithTimeout(context.Background(), stopTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("server forced to shutdown", zap.Error(err))
	}
}
