package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

// Server exposes RESTful API endpoints.
type Server struct {
	Router *httprouter.Router
	logger *log.Logger
}

// New instantiates an HTTP server and builds a route table.
func New() *Server {
	s := Server{
		Router: httprouter.New(),
		logger: log.New(os.Stderr, "API: ", 0),
	}
	s.buildRouteTable()
	return &s
}

func (s *Server) buildRouteTable() {
	s.Router.GET("/liveness", s.handleLiveness())
}

// encodeJSON takes content of any type (v) and encodes to the writer (w) in
// JSON format.
func (s *Server) encodeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		Error(w, err, http.StatusInternalServerError)
	}
}

// handleLiveness acts as the Kubernetes liveness probe.
func (s *Server) handleLiveness() httprouter.Handle {
	type response struct {
		Message string `json:"message"`
	}
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		payload := response{Message: "Application healthy!"}
		s.encodeJSON(w, payload)
	}
}