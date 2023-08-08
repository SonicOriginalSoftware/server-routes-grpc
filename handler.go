//revive:disable:package-comments

package grpc

import (
	"net/http"
	"os"

	"git.sonicoriginal.software/logger.git"
	"git.sonicoriginal.software/server.git/v2"
)

// Name is the name used to identify the service
const name = "grpc"

// Handler handles GRPC API requests
type handler struct {
	logger logger.Log
}

// ServeHTTP fulfills the http.Handler contract for Handler
func (handler *handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	http.Error(writer, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

// New returns a new Handler
func New(mux *http.ServeMux) (route string) {
	logger := logger.New(
		name,
		logger.DefaultSeverity,
		os.Stdout,
		os.Stderr,
	)

	return server.RegisterHandler(name, &handler{logger}, mux)
}
