package log

import (
	"github.com/stretchr/testify/assert"
	"go-cheatsheet/log"
	"testing"
	"time"
)

func TestGetElapsedTime(t *testing.T) {
	t.Run("500 millisecond", func(t *testing.T) {
		// given
		duration500Millisecond := 500 * time.Millisecond

		// when
		elapsedTime := log.GetElapsedTime(func() {
			time.Sleep(duration500Millisecond)
		})

		// then
		assert.True(t, elapsedTime >= duration500Millisecond)
	})

}
