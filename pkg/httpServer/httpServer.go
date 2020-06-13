package httpServer

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/smatton/httpServerTemplate/pkg/http/handler"
	"github.com/smatton/httpServerTemplate/pkg/http/webserver"
)

type Config struct {
	Port   string
	Exit   chan os.Signal
	Done   chan bool
	Logger *log.Logger
	Server *http.Server
	Router *http.ServeMux
}

func New(port string) *Config {
	var cfg Config
	// Initialize logger
	cfg.Logger = log.New(os.Stdout, "INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	done := make(chan bool, 1)
	exit := make(chan os.Signal, 1)
	cfg.Done = done
	cfg.Exit = exit
	cfg.Port = port

	// Create new simple server
	cfg.Server, cfg.Router = webserver.NewSimpleServer(cfg.Logger, port)

	// minimally add the alive handle

	cfg.Router.HandleFunc("/alive", func(w http.ResponseWriter, r *http.Request) {
		handler.Alive(w, r)
	})
	cfg.Router.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		handler.ShutDown(w, r, cfg.Exit)
	})

	return &cfg
}

func (cfg *Config) Start() error {

	// start Gracefull shutdown thread
	signal.Notify(cfg.Exit, os.Interrupt)
	go webserver.GracefullShutdown(cfg.Server, cfg.Logger, cfg.Exit, cfg.Done)
	cfg.Logger.Println("Listening on: ", ":"+cfg.Port)

	if err := cfg.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		cfg.Logger.Fatalf("Could not listen on %s: %v\n", ":"+cfg.Port, err)
		return err
	}

	<-cfg.Done
	return nil
}
