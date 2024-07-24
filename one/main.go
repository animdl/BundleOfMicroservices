package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	// creates the service stack to be LoggingService -> CatFactService
	service := NewCatFactService("https://catfact.ninja/fact")
	service = NewLoggingService(service)

	// calls the receiver function of the service
	// this returns the fact from the url provided to the service
	fact, err := service.GetCatFact(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	// prints the fact returned from the service
	fmt.Printf("%+v\n", fact)
}
