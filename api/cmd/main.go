package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/igormichalak/portfolio/api/internal/models"
)

const (
	DefaultAPIPort = 4000
	DefaultDSN     = "postgres://web:1234@db:5432/portfolio"
)

type config struct {
	port int
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

	flag.IntVar(&cfg.port, "port", 0, "API server port")
	flag.StringVar(&cfg.dsn, "dsn", "", "PostgreSQL data source name")

	flag.Parse()

	loadEnvInt(&cfg.port, "API_PORT", DefaultAPIPort, false)
	loadEnvString(&cfg.dsn, "DSN", DefaultDSN, false)

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	db, err := openDB(cfg.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		config:    cfg,
		errorLog:  errorLog,
		infoLog:   infoLog,
		blogPosts: &models.BlogPostModel{DB: db},
		blogTags:  &models.BlogTagModel{DB: db},
	}

	err = app.serve()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(context.Background()); err != nil {
		return nil, err
	}
	return db, nil
}
