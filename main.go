package main

import (
	"log"
	"net/http"

	"github.com/atomist/gokit-service-demo/gokitdemo"
)

func main() {
	gokitdemo.AddServices()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
