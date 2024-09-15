package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/K-Kizuku/eisa-auth/internal/di"
	env "github.com/K-Kizuku/eisa-auth/pkg/config"
	"github.com/K-Kizuku/eisa-auth/pkg/handler"
)

func main() {
	env.LoadEnv()
	ctx := context.Background()
	h := di.InitHandler()
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello eisa")
		w.WriteHeader(http.StatusOK)
	})

	mux.Handle("POST /signup", handler.AppHandler(h.UserHandler.SignUp()))
	mux.Handle("POST /signin", handler.AppHandler(h.UserHandler.SignIn()))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			slog.Error("server error", "error", err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("server error", "error", err.Error())
	}
}
