package log

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

var Logger = NewDefaultLogger()

func NewDefaultLogger() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	return zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime}).With().Timestamp().Logger()
}

func SetGlobalLogLevel(level string) {
	if !setGlobalLogLevel(level) {
		Logger.Info().Msg("the error level string did not match an error level")
	}
}

func setGlobalLogLevel(level string) bool {
	if lvl, err := zerolog.ParseLevel(level); err != nil || lvl == zerolog.NoLevel {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		return false
	} else {
		zerolog.SetGlobalLevel(lvl)
		return true
	}
}

func Err(err error) *zerolog.Event {
	return Logger.Err(err)
}

func Trace() *zerolog.Event {
	return Logger.Trace()
}

func Debug() *zerolog.Event {
	return Logger.Debug()
}

func Info() *zerolog.Event {
	return Logger.Info()
}

func Warn() *zerolog.Event {
	return Logger.Warn()
}

func Error() *zerolog.Event {
	return Logger.Error()
}

func Fatal() *zerolog.Event {
	return Logger.Fatal()
}

func Panic() *zerolog.Event {
	return Logger.Panic()
}
