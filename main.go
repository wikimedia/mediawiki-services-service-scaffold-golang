package main

import  "net/http"

type server struct{}

func main() {
	healthzHandler := HealthZ{}
    http.Handle("/health", &healthzHandler)
    http.ListenAndServe(":8000", nil)
}