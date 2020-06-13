package handler

import (
	"fmt"
	"net/http"
	"os"
)

//Alive returns to 200
func Alive(writer http.ResponseWriter, request *http.Request) {

	writer.WriteHeader(http.StatusOK)
}

func GracefullShutdown(writer http.ResponseWriter, request *http.Request, quit chan<- os.Signal) {

	writer.WriteHeader(http.StatusOK)
	quit <- os.Signal(os.Interrupt)

}

func ShutDown(writer http.ResponseWriter, request *http.Request, shutdown chan<- os.Signal) {
	writer.WriteHeader(http.StatusOK)
	fmt.Fprint(writer, "Server Shutdown")
	close(shutdown)

}
