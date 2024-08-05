package logging

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/saphieron/gowatcher/liberrors"
)

var programLevel = new(slog.LevelVar) // Info by default, kept in case i want to switch the log level at runtime

func Init(logLevel uint8) error {
	var output io.Writer = os.Stderr
	switch logLevel {
	case 0:
		output = io.Discard
	case 1:
		programLevel.Set(slog.LevelError)
	case 2:
		programLevel.Set(slog.LevelWarn)
	case 3:
		programLevel.Set(slog.LevelInfo)
	case 4:
		programLevel.Set(slog.LevelDebug)
	default:
		return liberrors.NewInvalidArgsError(fmt.Sprintf("got invalid log level value '%d'", logLevel))
	}

	logger := slog.New(slog.NewTextHandler(output, &slog.HandlerOptions{
		Level: programLevel,
	}))
	slog.SetDefault(logger)
	return nil
}
