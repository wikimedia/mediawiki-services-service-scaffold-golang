package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"
)

var (
	version   = "unknown"
	buildDate = "unknown"
	buildHost = "unknown"
)

// Entrypoint for our service
func main() {
	var confFile = flag.String("config", "./config.yaml", "Path to the configuration file")
	fmt.Println(confFile)
	var config *Config
	var err error

	flag.Parse()

	if config, err = ReadConfig(*confFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println((config))
	fmt.Printf("Initializing service %s (Go version: %s, Build host: %s, Timestamp: %s", config.ServiceName, version, buildHost, buildDate)

	fmt.Println(err)
	if err != nil {
		// log.Fatal(err)
		// ensure this is stderr
		// exit with a non zero status
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// logger, err := NewLogger(os.Stdout, config.ServiceName, config.LogLevel)
	fmt.Println(path.Join(config.BaseURI, "healthz"))
	// logger.Info("Initializing Kask %s (Go version: %s, Build host: %s, Timestamp: %s)...", version, runtime.Version(), buildHost, buildDate)
	http.Handle(path.Join(config.BaseURI, "healthz"), &HealthzHandler{NewHealthz(version, buildDate, buildHost)})
	http.Handle(path.Join(config.BaseURI, "echo"), &EchoHandler{})
	http.ListenAndServe(fmt.Sprintf("%s:%d", config.Address, config.Port), nil)
}
