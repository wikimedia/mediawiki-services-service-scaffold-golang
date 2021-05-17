package main
import "net/http"

type HealthZ struct {
	// get version, host, date, etc. from Makefile
}


func (s *HealthZ) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "hello world!"}`))
}

