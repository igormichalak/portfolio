package main

import (
	"flag"
	"log"
	"os"
)

type config struct {
	port uint
}

type application struct {
	config   config
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	var cfg config

	flag.UintVar(&cfg.port, "port", 4000, "API server port")

	flag.Parse()

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := &application{
		config:   cfg,
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	err := app.serve()
	errorLog.Fatal(err)
}
