package main

import (
	"encoding/json"
	"net/http"
	"runtime"
)

// Healthz is a struct returned by /healthz endpoint
type Healthz struct {
	Version   string `json:"version"`
	BuildDate string `json:"build_date"`
	BuildHost string `json:"build_host"`
	GoVersion string `json:"go_version"`
}

// NewHealthz creates and returns a new HealthZ ...
func NewHealthz(version, date, host string) *Healthz {
	return &Healthz{
		Version:   version,
		BuildDate: date,
		BuildHost: host,
		GoVersion: runtime.Version(),
	}
}

type HealthzHandler struct {
	healthz *Healthz
}

// Serves Healthz struct data as JSON
func (s *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var response []byte
	var err error

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if response, err = json.MarshalIndent(s.healthz, "", "  "); err != nil {
		w.Write([]byte(`{}`))
		return
	}
	w.Write(response)
}
