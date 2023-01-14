package grpc_test

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"git.sonicoriginal.software/routes/grpc"
	lib "git.sonicoriginal.software/server"
)

var certs []tls.Certificate

func TestHandler(t *testing.T) {
	route := fmt.Sprintf("localhost/%v/", grpc.Name)
	t.Setenv(fmt.Sprintf("%v_SERVE_ADDRESS", strings.ToUpper(grpc.Name)), route)

	grpc.New()

	ctx, cancelFunction := context.WithCancel(context.Background())
	address, errChan := lib.Run(ctx, certs)

	// TODO modify the request to send a proper grpc request
	url := fmt.Sprintf("http://%v/%v/", address, grpc.Name)
	response, err := http.DefaultClient.Get(url)
	if err != nil {
		t.Fatalf("%v\n", err)
	}

	cancelFunction()

	if err := <-errChan; err != nil {
		t.Fatalf("Server errored: %v", err)
	}

	if response.Status != http.StatusText(http.StatusNotImplemented) && response.StatusCode != http.StatusNotImplemented {
		t.Fatalf("Server returned: %v", response.Status)
	}
}
