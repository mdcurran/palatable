package server

import (
    "encoding/json"
    "net/http"

    "github.com/julienschmidt/httprouter"
    "go.uber.org/zap"
)

// Server exposes RESTful API endpoints.
type Server struct {
    Router *httprouter.Router
    Logger *zap.SugaredLogger
}

// New instantiates an HTTP server and builds a route table.
func New() *Server {
    s := Server{
        Router: httprouter.New(),
        Logger: initialiseLogger(),
    }
    s.buildRouteTable()
    return &s
}

func initialiseLogger() *zap.SugaredLogger {
    logger, err := zap.NewProduction()
    if err != nil {
        panic(err)
    }
    // Use the zap.SugaredLogger as recommended in the zap documentation:
    // https://godoc.org/go.uber.org/zap#hdr-Choosing_a_Logger
    return logger.Sugar()
}

func (s *Server) buildRouteTable() {
    s.Router.GET("/liveness", s.handleLiveness)
    s.Router.POST("/restaurant", s.postRestaurant)
    s.Router.POST("/item", s.postItem)
    s.Router.GET("/restaurant/:name", s.getRestaurant)
    s.Router.GET("/restaurant/:name/items", s.listItems)
}

// encodeJSON takes content of any type (v) and encodes to the writer (w) in
// JSON format.
func (s *Server) encodeJSON(w http.ResponseWriter, v interface{}) {
    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(v)
    if err != nil {
        Error(w, err, http.StatusInternalServerError)
        return
    }
}

// handleLiveness acts as the Kubernetes liveness probe.
func (s *Server) handleLiveness(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    type response struct {
        Message string `json:"message"`
    }
    payload := response{Message: "Application healthy!"}
    s.encodeJSON(w, payload)
}
