package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	h "com.github/imperium/internal/app/imperium/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.NotFound(h.NewNotFoundHandler().ServeHTTP)
		r.Get("/", h.NewHomeHandler().ServeHTTP)
	})

	// os kill commands and server start

	killsig := make(chan os.Signal, 1)

	signal.Notify(killsig, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    ":3333",
		Handler: r,
	}

	go func() {
		err := srv.ListenAndServe()

		if errors.Is(err, http.ErrServerClosed) {
			logger.Info("Server shutdown complete")
		} else if err != nil {
			logger.Error("Server error", slog.Any("err", err))
			os.Exit(1)
		}
	}()

	logger.Info("Server started", slog.String("port", "3333"), slog.String("env", "PROD"))
	<-killsig

	logger.Info("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Derver shutdown failed", slog.Any("err", err))
		os.Exit(1)
	}

	logger.Info("Derver shutdown complete")
}
