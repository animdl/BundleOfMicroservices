package main

import (
	"context"
	"encoding/json"
	"net/http"
)

// Service interface
// implements GetCatFact function
type Service interface {
	// argument - context
	// returns response in the structure of a CatFact, and an error
	GetCatFact(context.Context) (*CatFact, error)
}

// CatFactService structure
// contains the request url
type CatFactService struct {
	url string
}

// argument - url
// returns - a Service
// in this case, we return the CatFactService struct which implements the Service interface
func NewCatFactService(url string) Service {
	return &CatFactService{
		url: url,
	}
}

// receiver function from the Service interface for the CatFactService struct
// function that retrieves the fact itself
// argument - context
// uses - CatFactService struct
// returns (CatFact, error)
func (s *CatFactService) GetCatFact(c context.Context) (*CatFact, error) {
	// gets the response from the url
	response, err := http.Get(s.url)
	// if there is an error in the process, return the error
	if err != nil {
		return nil, err
	}

	// defer places the function after it at the end of the stack
	// executes the function right before the parent returns
	// used to clean up open resources
	defer response.Body.Close()

	/* Format Checking */
	// set fact to the structure of a CatFact
	fact := &CatFact{}
	// json.NewDecoder returns a Decoder with the data from the argument
	// .Decode reads the data from the Decoder and stores it to the argument
	// if there is an error in the process, return the error
	if err := json.NewDecoder(response.Body).Decode(fact); err != nil {
		return nil, err
	}

	// returns the fact
	return fact, nil
}
