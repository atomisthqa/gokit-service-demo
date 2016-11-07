package gokitdemo

import (
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

type gokitdemocontract interface {
	HelloWorld() (string, error)
}

type gokitdemo struct{}

// AddServices sets up and starts the service.
func AddServices() {
	ctx := context.Background()
	svc := gokitdemo{}

	gokitdemoHandler := httptransport.NewServer(
		ctx,
		makeHelloWorldEndpoint(svc),
		decodeHelloWorldRequest,
		encodeResponse,
	)

	http.Handle("/", gokitdemoHandler)
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
