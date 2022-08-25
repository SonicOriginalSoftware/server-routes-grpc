//revive:disable:package-comments

package grpc

import (
	"server/env"
	"server/logging"
	"server/net/local"

	"fmt"
	"net/http"
)

const prefix = "grpc"

// Handler handles GRPC API requests
type Handler struct {
	logger *logging.Logger
}

// ServeHTTP fulfills the http.Handler contract for Handler
func (handler *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	http.Error(writer, "Not yet implemented!", http.StatusNotImplemented)
}

// Prefix is the subdomain prefix
func (handler *Handler) Prefix() string {
	return prefix
}

// Address returns the address the Handler will service
func (handler *Handler) Address() string {
	return env.Address(prefix, fmt.Sprintf("%v.%v", prefix, local.Path("")))
}

// New returns a new Handler
func New() *Handler {
	return &Handler{
		logger: logging.New(prefix),
	}
}
