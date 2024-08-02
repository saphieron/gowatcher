package logging

import (
	"log/slog"
	"os"
)

var programLevel = new(slog.LevelVar) // Info by default

func Init(verbose bool) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	slog.SetDefault(logger)
	if verbose {
		programLevel.Set(slog.LevelDebug)
	} else {
		programLevel.Set(slog.LevelWarn)
	}
}
