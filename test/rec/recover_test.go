package rec

import (
	"github.com/stretchr/testify/assert"
	"go-cheatsheet/rec"
	"sync"
	"testing"
)

func TestSyncWithRecover(t *testing.T) {
	// given
	runnable := func() error {
		panic("test_panic")
	}

	// when
	err := rec.SyncWithRecover(runnable, rec.DefaultSyncRecoverFunc)

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
	rec.AsyncWithRecover(runnable, func(rec rec.StackError) {
		assert.Equal(t, "test_panic", rec.Error())
		wg.Done()
	})
	wg.Wait()

	// then
	assert.True(t, true, true)
}
