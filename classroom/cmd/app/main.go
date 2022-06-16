package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/lucasd-coder/classroom/internal/app"
	"github.com/lucasd-coder/classroom/internal/pkg/logger"
)

const defaultPort = "3334"

func main() {
	errs := godotenv.Load()
	if errs != nil {
		logger.Log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	app.Run(port)
}
