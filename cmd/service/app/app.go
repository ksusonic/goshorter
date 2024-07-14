package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ksusonic/goshorter/internal/repository"
	"github.com/ksusonic/goshorter/internal/service/shortener"
)

func Run() {
	var (
		ctx    = context.Background()
		cfg    = newConfig()
		logger = log.Default()
	)

	repo, closer := repository.NewRepository(ctx, cfg.DatabaseDSN)
	defer closer()

	shortenerService := shortener.NewService(cfg.ShortURLPrefix, repo, logger)

	r := setupRouter(shortenerService)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      r.Handler(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("srv.Shutdown: %+v", err)
		return
	}
}
