package pool

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-cheatsheet/pool"
	"go-cheatsheet/utils"
	"sync"
	"testing"
)

func TestAsyncPool(t *testing.T) {
	// given
	var wg sync.WaitGroup
	runner := pool.NewAsyncPool(context.Background(), func(option *pool.AsyncPoolOption) {
		option.BatchSize = utils.Pointer(10)
	})
	var arr []int
	testSize := 2000
	asyncSum := 0
	expectedSum := testSize * (testSize + 1) / 2

	type testStruct struct {
		value int
	}

	// when
	wg.Add(testSize)
	for i := 0; i < testSize; i++ {
		obj := testStruct{
			value: i,
		}

		runner.SubmitJob(func(synchronizer pool.Synchronizer) {
			defer func() {
				wg.Done()
			}()

			synchronizer.Sync(func() {
				arr = append(arr, obj.value+1)
				asyncSum = asyncSum + obj.value + 1
			})
		})
	}
	wg.Wait()

	assert.Equal(t, expectedSum, asyncSum)
	assert.True(t, len(arr) == testSize)

	asyncSum = 0
	arr = []int{}
	wg.Add(testSize)
	for i := 0; i < testSize; i++ {
		j := i

		runner.SubmitJob(func(synchronizer pool.Synchronizer) {
			defer func() {
				wg.Done()
			}()

			synchronizer.Sync(func() {
				arr = append(arr, j+1)
				asyncSum = asyncSum + j + 1
			})
		})
	}
	wg.Wait()

	assert.Equal(t, expectedSum, asyncSum)
	assert.True(t, len(arr) == testSize)
}

func TestQuiteAsyncRunner(t *testing.T) {
	// given
	var wg sync.WaitGroup
	ctx, cancelFunc := context.WithCancel(context.Background())
	runner := pool.NewAsyncPool(ctx, func(option *pool.AsyncPoolOption) {
		option.BatchSize = utils.Pointer(10)
	})
	var arr []int
	testSize := 2000
	asyncSum := 0
	expectedSum := testSize * (testSize + 1) / 2

	type testStruct struct {
		value int
	}

	// when
	wg.Add(testSize)
	for i := 0; i < testSize; i++ {
		obj := testStruct{
			value: i,
		}

		runner.SubmitJob(func(synchronizer pool.Synchronizer) {
			defer func() {
				wg.Done()
			}()

			synchronizer.Sync(func() {
				arr = append(arr, obj.value+1)
				asyncSum = asyncSum + obj.value + 1
			})
		})
	}
	wg.Wait()

	assert.Equal(t, expectedSum, asyncSum)
	assert.True(t, len(arr) == testSize)

	asyncSum = 0
	arr = []int{}
	wg.Add(testSize)
	for i := 0; i < testSize; i++ {
		j := i

		runner.SubmitJob(func(synchronizer pool.Synchronizer) {
			defer func() {
				wg.Done()
			}()

			synchronizer.Sync(func() {
				arr = append(arr, j+1)
				asyncSum = asyncSum + j + 1
			})
		})
	}
	wg.Wait()

	assert.Equal(t, expectedSum, asyncSum)
	assert.True(t, len(arr) == testSize)

	cancelFunc()
}
