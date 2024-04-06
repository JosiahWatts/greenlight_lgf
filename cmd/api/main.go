package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// later we'll generate this at build time
const version = "1.0.0"

// define a config struct to hold config settings
type config struct {
	port int
	env  string
}

// holds dependencies for our http handlers, helpers, middlewares.
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	// read the value of the port and env command-line flags into the config struct.
	// We default to using port 4000 and 'development' env if no flags are provided
	// https://pkg.go.dev/flag
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// initialize new logger which writes messages to std out stream,
	// prefixed with current date and time
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// declare an instance of app struct
	app := &application{
		config: cfg,
		logger: logger,
	}

	// declare a http server with some timeout settings
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// start the server
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
