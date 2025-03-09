package log

import (
	"go-cheatsheet/gen"
	"log/slog"
	"os"
)

var defaultLogger = gen.NewSingletonSupplier(func() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	}))
})

func GetDefaultLogger() *slog.Logger {
	return defaultLogger.Get()
}
