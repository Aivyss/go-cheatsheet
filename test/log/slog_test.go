package log

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-cheatsheet/log"
	"log/slog"
	"testing"
)

func TestGetDefaultLogger(t *testing.T) {
	t.Run("GetDefaultLogger", func(t *testing.T) {
		// given
		ctx := context.Background()
		// when 1
		firstLogger := log.GetDefaultLogger()

		// then 1
		assert.NotNil(t, firstLogger)
		assert.False(t, firstLogger.Enabled(ctx, slog.LevelDebug))
		assert.True(t, firstLogger.Enabled(ctx, slog.LevelInfo))
		assert.True(t, firstLogger.Enabled(ctx, slog.LevelWarn))
		assert.True(t, firstLogger.Enabled(ctx, slog.LevelError))

		// when 2
		secondLogger := log.GetDefaultLogger()

		// then 2
		assert.Equal(t, firstLogger, secondLogger)
	})
}
