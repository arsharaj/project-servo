package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init() {
	// Set global log time format and level
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Set output to stderr in JSON format
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: zerolog.TimeFieldFormat,
	})
}
