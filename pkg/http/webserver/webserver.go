package webserver

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"
)

func GracefullShutdown(server *http.Server, logger *log.Logger, quit <-chan os.Signal, done chan<- bool) {
	<-quit
	logger.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
	close(done)
}

func NewSimpleServer(logger *log.Logger, port string) (*http.Server, *http.ServeMux) {
	router := http.NewServeMux()

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return server, router
}
