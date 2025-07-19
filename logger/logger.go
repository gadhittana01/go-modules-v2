package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Logger zerolog.Logger

func Init() {
	// Set log level from environment variable
	logLevel := strings.ToLower(os.Getenv("LOG_LEVEL"))
	if logLevel == "" {
		logLevel = "info"
	}

	var level zerolog.Level
	switch logLevel {
	case "debug":
		level = zerolog.DebugLevel
	case "info":
		level = zerolog.InfoLevel
	case "warn":
		level = zerolog.WarnLevel
	case "error":
		level = zerolog.ErrorLevel
	case "fatal":
		level = zerolog.FatalLevel
	case "panic":
		level = zerolog.PanicLevel
	default:
		level = zerolog.InfoLevel
	}

	// Configure zerolog
	zerolog.SetGlobalLevel(level)

	// Use pretty console writer for development
	if os.Getenv("ENV") == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	} else {
		// Use JSON format for production - output to stdout for better container log compatibility
		log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}

	Logger = log.Logger
}

// Convenience functions
func Debug(msg string, fields map[string]interface{}) {
	if fields != nil {
		Logger.Debug().Fields(fields).Msg(msg)
	} else {
		Logger.Debug().Msg(msg)
	}
}

func Info(msg string, fields map[string]interface{}) {
	if fields != nil {
		Logger.Info().Fields(fields).Msg(msg)
	} else {
		Logger.Info().Msg(msg)
	}
}

func Warn(msg string, fields map[string]interface{}) {
	if fields != nil {
		Logger.Warn().Fields(fields).Msg(msg)
	} else {
		Logger.Warn().Msg(msg)
	}
}

func Error(msg string, err error, fields map[string]interface{}) {
	if fields != nil {
		Logger.Error().Err(err).Fields(fields).Msg(msg)
	} else {
		Logger.Error().Err(err).Msg(msg)
	}
}

func Fatal(msg string, err error, fields map[string]interface{}) {
	if fields != nil {
		Logger.Fatal().Err(err).Fields(fields).Msg(msg)
	} else {
		Logger.Fatal().Err(err).Msg(msg)
	}
}
