package main

import (
	"log"
	"net/http"
)

const webPort = ":8080"

type Config struct{}


func main() {
	app := Config{}
	log.Printf("Starting Broker Service on port %s", webPort)
	srv := &http.Server{
		Addr:    webPort,
		Handler: app.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
