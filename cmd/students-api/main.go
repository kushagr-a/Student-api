package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"student_api/internal/config"
	"student_api/internal/http/handlers/student"
	"syscall"
	"time"
)

func main() {
	// load config

	cfg := config.MustLoad()

	// database setup
	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())
	// setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Println("Server Started at", cfg.Addr)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// start server
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("error starting server: %s", err.Error())
		}
	}()

	<-done

	slog.Info("Shutting down the server")

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("error shutting down the server: %s", err.Error())
	}

	slog.Info("server Shutdown successfully")
}
