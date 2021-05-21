package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"

	log "github.com/eevans/servicelib-golang/logger"
)

var (
	version   = "unknown"
	buildDate = "unknown"
	buildHost = "unknown"
)

// Entrypoint for our service
func main() {
	var confFile = flag.String("config", "./config.yaml", "Path to the configuration file")

	var config *Config
	var err error
	var logger *log.Logger

	flag.Parse()

	if config, err = ReadConfig(*confFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	logger, err = log.NewLogger(
		os.Stdout,
		config.ServiceName,
		config.ServiceType,
		log.INFO, // TODO: Change to use config
	)

	if err != nil {
		logger.Error("An error occurred:", err)
	}

	logger.Info("Initializing service %s (Go version: %s, Build host: %s, Timestamp: %s", config.ServiceName, version, buildHost, buildDate)

	http.Handle("healthz", &HealthzHandler{NewHealthz(version, buildDate, buildHost)})
	http.Handle(path.Join(config.BaseURI, "echo"), &EchoHandler{})
	http.ListenAndServe(fmt.Sprintf("%s:%d", config.Address, config.Port), nil)
}
