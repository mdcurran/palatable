package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Error produces a JSON-encoded application error with an appropriate HTTP
// status code.
func Error(w http.ResponseWriter, err error, code int) error {
	e := errorObject{Detail: err.Error(), Status: code}
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(&e)
}

// errorObject is a structured error providing information about an application
// problem.
type errorObject struct {
	Detail string `json:"detail"`
	Status int    `json:"status"`
}

func (e *errorObject) Error() string {
	return fmt.Sprintf("Error: %s Status Code: %d\n", e.Detail, e.Status)
}
