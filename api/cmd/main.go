package main

import (
	"log"
)

type config struct {
	port int
}

type application struct {
	config config
}

func main() {
	var cfg config
	cfg.port = 4000

	app := &application{
		config: cfg,
	}

	err := app.serve()
	log.Fatal(err)
}
