package server

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestServer_HandleLiveness(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/api/liveness", nil)
    res := httptest.NewRecorder()

    s := New()

    s.Router.ServeHTTP(res, req)

    if res.Code != http.StatusOK {
        t.Fatalf("got %d, want %d", res.Code, http.StatusOK)
    }

    body := strings.Trim(res.Body.String(), "\n")
    want := `{"message":"Application healthy!"}`
    if body != want {
        t.Errorf("got %q, want %q", body, want)
    }
}
