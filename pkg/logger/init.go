package logger

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init() {
	// Set logger output stream and time format
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal().AnErr("Error loading .env file: %w", err)
	}
}
