//revive:disable:package-comments

package grpc

import (
	"git.nathanblair.rocks/server/logging"

	"net/http"
)

// Prefix is the name used to identify the service
const Prefix = "grpc"

// Handler handles GRPC API requests
type Handler struct {
	logger *logging.Logger
}

// ServeHTTP fulfills the http.Handler contract for Handler
func (handler *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	http.Error(writer, "Not yet implemented!", http.StatusNotImplemented)
}

// New returns a new Handler
func New() *Handler {
	return &Handler{
		logger: logging.New(Prefix),
	}
}
