package main

import (
	"context"
	"encoding/json"
	"net/http"
)

// ApiServer structure
type ApiServer struct {
	service Service
}

// argument - a service
// returns - an ApiServer struct
func NewApiServer(service Service) *ApiServer {
	return &ApiServer{
		service: service,
	}
}

// argument - port number
// returns - json response
// starts a server on the "/" route for localhost at the given port
// calls the handleGetCatFact function on requests
func (s *ApiServer) Start(listenAddr string) error {
	http.HandleFunc("/", s.handleGetCatFact)
	return http.ListenAndServe(listenAddr, nil)
}

// receiver function for the ApiServer struct
// after a request is made, a fact is retrieved and a json response is created
func (s *ApiServer) handleGetCatFact(w http.ResponseWriter, r *http.Request) {
	// gets the response from the GetCatFact function
	fact, err := s.service.GetCatFact(context.Background())
	// if there is an error in retrieving the response, return an error HTTP response
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, fact)
}

// argument - http.ResponseWriter, http code, any data
// returns - json response
// creates a json for http responses
func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
