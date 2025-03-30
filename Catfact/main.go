package main

import (
	"log"
)

func main() {
	// creates the service stack to be LoggingService -> CatFactService
	service := NewCatFactService("https://catfact.ninja/fact")
	service = NewLoggingService(service)

	// assigns an api server a service to run
	// starts the server on the given port
	apiServer := NewApiServer(service)
	log.Fatal(apiServer.Start(":3000"))
}
