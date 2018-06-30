package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/yadunut/sma-website/http"
	"github.com/yadunut/sma-website/postgresql"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	var (
		port      = os.Getenv("PORT")
		logs      = os.Getenv("LOGS")
		dbConnStr = os.Getenv("DB")
		debug     = os.Getenv("DEBUG")
	)

	logger := logrus.New()
	if debug != "" {
		logger.Formatter = &logrus.TextFormatter{ForceColors: true}
		logger.Level = logrus.DebugLevel
	}

	if logs != "" {
		f, err := os.Open(logs)
		if err != nil {
			panic(err)
		}

		logger.Out = f
	} else {
		logger.Out = os.Stdout

		if port == "" {
			port = ":8080"
		}

		db, err := postgresql.Open(dbConnStr)
		if err != nil {
			logger.Error(err)
		}
		defer db.Close()

		server := http.Server{
			Port:     port,
			Logger:   logger,
			Database: db,
		}

		err = server.Listen()
		if err != nil {
			logger.Error(err)
		}

		os.Exit(1)
	}
}
