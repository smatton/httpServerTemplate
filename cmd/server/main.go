package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/smatton/httpServerTemplate/pkg/http/handler"
	"github.com/smatton/httpServerTemplate/pkg/http/webserver"
	"github.com/smatton/httpServerTemplate/pkg/network"
)

var (
	PORT string
)

func main() {

	flag.StringVar(&PORT, "port", "9023", "port to start registry on")

	flag.Parse()
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	myIP, err := network.GetOutboundIP()
	if err != nil {
		logger.Println("Couldn't determine hostname, starting on loopback 127.0.0.1")
		myIP = "127.0.0.1"
	}

	// channels for gracefully shut down
	done := make(chan bool, 1)
	exit := make(chan os.Signal, 1)

	signal.Notify(exit, os.Interrupt)
	server, router := webserver.NewSimpleServer(logger, PORT)

	// Set up routes
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// function for handling im here from servers
		handler.Ping(w, r)
	})
	logger.Println("Listening on ", myIP+":"+PORT)

	go webserver.GracefullShutdown(server, logger, exit, done)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", ":"+PORT, err)
	}
	<-done
	logger.Println("Server stopped")
}
