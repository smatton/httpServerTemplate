package handler

import (
	"net/http"
)

//Ping returns to 200 to caller
func Ping(writer http.ResponseWriter, request *http.Request) {

	writer.WriteHeader(http.StatusOK)
}
