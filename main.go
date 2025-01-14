package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sivchari/gosched-simulator/handler"
)

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
}

// flush added goroutine
// add goroutine -> new proc with goroutine name

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	mux := http.NewServeMux()
	mux.HandleFunc("/g", handler.G())
	mux.HandleFunc("/m", handler.M())
	mux.HandleFunc("/p", handler.P())
	mux.HandleFunc("/sched", handler.Sched())

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		slog.Info("starting server")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	<-ctx.Done()

	slog.Info("shutting down server")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		panic(err)
	}

	slog.Info("server shutdown")
}
