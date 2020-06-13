package main

import (
	"flag"
	"log"
	"os"

	"github.com/smatton/httpServerTemplate/pkg/httpServer"
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

	logger.Println("Listening on ", myIP+":"+PORT)
	myserver := httpServer.New(PORT)

	myserver.Start()

}
