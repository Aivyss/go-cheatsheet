package panic_recover

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestSyncWithRecover(t *testing.T) {
	// given
	runnable := func() error {
		panic("test_panic")
	}

	// when
	err := SyncWithRecover(runnable, DefaultSyncRecoverFunc)

	// then
	assert.Equal(t, "test_panic", err.Error())
}

func TestAsyncWithRecoverFunc(t *testing.T) {
	// given
	runnable := func() error {
		panic("test_panic")
	}

	// when
	var wg sync.WaitGroup
	wg.Add(1)
	AsyncWithRecover(runnable, func(rec StackError) {
		assert.Equal(t, "test_panic", rec.Error())
		wg.Done()
	})
	wg.Wait()

	// then
	assert.True(t, true, true)
}
