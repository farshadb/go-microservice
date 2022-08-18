package main

import (
	"fmt"
	"log"
	"net/http"
	"microservice/broker-service/cmd/api/routes"
)

const webPort = "8000"

type Config struct {}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s\n", webPort)

	// define http server

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}