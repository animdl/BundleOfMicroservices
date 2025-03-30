package main

import (
	"context"
	"fmt"
	"time"
)

/* type of middleware to log service calls */

// logging service structure
// contains the service to log
type LoggingService struct {
	next Service
}

// argument - a Service
// returns - a Service
// returns a LoggingService struct which implements the Service interface
func NewLoggingService(next Service) Service {
	return &LoggingService{
		next: next,
	}
}

// receiver function from the Service interface for the LoggingService struct
// argument - context
// uses - LoggingService struct
// returns (CatFact, error)
func (s *LoggingService) GetCatFact(c context.Context) (fact *CatFact, err error) {

	// defering an anonymous function that has the start time as an argument
	// function prints the fact and error from the named returns
	// and the time taken to start the service and return a response
	defer func(start time.Time) {
		fmt.Printf("fact:%v \nerr:%v \ntime:%v \n", fact.Fact, err, time.Since(start))
	}(time.Now())

	// returns the response from the GetCatFact function
	return s.next.GetCatFact(c)
}
