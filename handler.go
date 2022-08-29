//revive:disable:package-comments

package grpc

import (
	"git.nathanblair.rocks/server/handlers"
	"git.nathanblair.rocks/server/logging"

	"net/http"
)

// Name is the name used to identify the service
const Name = "grpc"

// Handler handles GRPC API requests
type Handler struct {
	logger *logging.Logger
}

// ServeHTTP fulfills the http.Handler contract for Handler
func (handler *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	http.Error(writer, "Not yet implemented!", http.StatusNotImplemented)
}

// New returns a new Handler
func New() (handler *Handler) {
	logger := logging.New(Name)
	handler = &Handler{logger}
	handlers.Register(Name, handler, logger)

	return
}
