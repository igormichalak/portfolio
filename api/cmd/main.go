package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type config struct {
	port uint
	dsn  string
}

type application struct {
	config   config
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	var cfg config

	flag.UintVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.dsn, "dsn", "postgres://web:1234@db:5432/portfolio", "PostgreSQL data source name")

	flag.Parse()

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	dbpool, err := openDB(cfg.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer dbpool.Close()

	app := &application{
		config:   cfg,
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	err = app.serve()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	if err = dbpool.Ping(context.Background()); err != nil {
		return nil, err
	}
	return dbpool, nil
}
