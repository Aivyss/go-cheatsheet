package types

import (
	"github.com/stretchr/testify/assert"
	"go-cheatsheet/types"
	"testing"
)

func TestPair(t *testing.T) {
	t.Run("Pair", func(t *testing.T) {
		// given
		pair := types.PairOf(1, 2)

		// then
		assert.Equal(t, 1, pair.First())
		assert.Equal(t, 2, pair.Second())
	})
}
