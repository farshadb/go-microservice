package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "8000"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting front end service on port %s\n", webPort)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
