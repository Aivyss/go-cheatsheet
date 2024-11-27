package gen

import (
	"github.com/stretchr/testify/assert"
	"go-cheatsheet/gen"
	"sync"
	"testing"
	"time"
)

type singletonTestInterface interface {
	Value() string
	Modify(value string)
}
type singletonTestStruct struct {
	value string
}

func (s *singletonTestStruct) Value() string {
	return s.value
}
func (s *singletonTestStruct) Modify(value string) {
	s.value = value
}

func TestSingletonSupplier(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		// given
		cnt := 0
		supplier := gen.NewSingletonSupplier(func() singletonTestStruct {
			cnt++
			return singletonTestStruct{
				value: "a",
			}
		})

		// when
		objFirst := supplier.Get()
		objSecond := supplier.Get()

		// then
		assert.Equal(t, "a", objFirst.value)
		assert.Equal(t, "a", objSecond.value)
		assert.Equal(t, objFirst, objSecond)
		assert.Equal(t, 1, cnt)
	})
	t.Run("interface", func(t *testing.T) {
		// given
		cnt := 0
		supplier := gen.NewSingletonSupplier(func() singletonTestInterface {
			cnt++
			return &singletonTestStruct{
				value: "a",
			}
		})

		// when
		objFirst := supplier.Get()
		objSecond := supplier.Get()

		// then
		assert.Equal(t, "a", objFirst.Value())
		assert.Equal(t, "a", objSecond.Value())
		objFirst.Modify("b")
		assert.Equal(t, objFirst, objSecond)
		assert.Equal(t, "b", objFirst.Value())
		assert.Equal(t, "b", objSecond.Value())
		assert.Equal(t, 1, cnt)
	})
	t.Run("primitive", func(t *testing.T) {
		// given
		cnt := 0
		supplier := gen.NewSingletonSupplier(func() string {
			cnt++
			return "a"
		})

		// when
		first := supplier.Get()
		second := supplier.Get()

		// then
		assert.Equal(t, "a", first)
		assert.Equal(t, "a", second)
		assert.Equal(t, 1, cnt)
	})
	t.Run("async test(check goroutine safe)", func(t *testing.T) {
		cnt := 0
		supplier := gen.NewSingletonSupplier(func() string {
			cnt++
			time.Sleep(50 * time.Millisecond)
			return "a"
		})

		var wg sync.WaitGroup

		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				_ = supplier.Get()
				wg.Done()
			}()
		}
		wg.Wait()

		assert.Equal(t, 1, cnt)
	})
}
