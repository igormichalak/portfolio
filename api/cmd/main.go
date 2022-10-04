package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/igormichalak/portfolio/api/internal/models"
)

type config struct {
	port uint
	dsn  string
}

type application struct {
	config    config
	errorLog  *log.Logger
	infoLog   *log.Logger
	blogPosts *models.BlogPostModel
	blogTags  *models.BlogTagModel
}

func main() {
	var cfg config

	flag.UintVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.dsn, "dsn", "postgres://web:1234@db:5432/portfolio", "PostgreSQL data source name")

	flag.Parse()

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	dbPool, err := openDB(cfg.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer dbPool.Close()

	app := &application{
		config:    cfg,
		errorLog:  errorLog,
		infoLog:   infoLog,
		blogPosts: &models.BlogPostModel{DBPool: dbPool},
		blogTags:  &models.BlogTagModel{DBPool: dbPool},
	}

	err = app.serve()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	if err = dbPool.Ping(context.Background()); err != nil {
		return nil, err
	}
	return dbPool, nil
}
