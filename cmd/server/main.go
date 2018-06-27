package main

import (
	"github.com/yadunut/sma-website/http"
	"github.com/joho/godotenv"

	"os"
	"mod/github.com/sirupsen/logrus@v1.0.5"
)

func init() {
	if err := godotenv.Load(); err != nil {
	}
}

// Dependency inversion
// create the dependencies and pass it to the function which needs it.
func main() {

	var (
		port = os.Getenv("PORT")
		logs = os.Getenv("LOGS")
	)

	logger := logrus.New()

	if logs != "" {
		f, err := os.Open(logs)
		if err != nil {
			panic(err)
		}

		logger.Out = f
	} else {
		logger.Out = os.Stdout
	}

	if port == "" {
		port = "8080"
	}

	server := http.Server{
		Port:   port,
		Logger: logger,
	}
	server.Listen()
}
